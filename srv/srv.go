package srv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
	"github.com/bcspragu/Bananagrama/pubsub"
)

const (
	DumpSize    = 3
	scaleFactor = 1
	// TODO: Add MaxPlayers back in
)

type Server struct {
	r    *rand.Rand
	ps   *pubsub.PubSub
	db   banana.DB
	dict banana.Dictionary

	// Standard channels
	gameUpdateChannel pubsub.Channel

	sync.RWMutex
	gameChannels   map[banana.GameID]pubsub.Channel
	playerChannels map[banana.PlayerID]pubsub.Channel
}

func New(r *rand.Rand, db banana.DB, dict banana.Dictionary) (*Server, error) {
	ps := pubsub.New()
	gameUpdateChannel, err := ps.NewChannel("GAME_UPDATES")
	if err != nil {
		return nil, fmt.Errorf("failed to create new games PubSub channel: %v", err)
	}

	return &Server{
		r:                 r,
		ps:                ps,
		db:                db,
		dict:              dict,
		gameUpdateChannel: gameUpdateChannel,
		gameChannels:      make(map[banana.GameID]pubsub.Channel),
		playerChannels:    make(map[banana.PlayerID]pubsub.Channel),
	}, nil
}

type update struct{}

func (s *Server) NewGame(ctx context.Context, req *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("must specify a game name")
	}

	bunch, err := banana.NewBunch(banana.Bananagrams(), scaleFactor)
	if err != nil {
		return nil, fmt.Errorf("failed to make bunch: %v", err)
	}

	id, err := s.db.NewGame(name, bunch)
	if err != nil {
		return nil, err
	}

	s.Lock()
	ch, err := s.ps.NewChannel(string(id))
	if err != nil {
		log.Printf("failed to create channel for game %q: %v", id, err)
		s.Unlock()
		return nil, fmt.Errorf("failed to create channel for game %q: %v", id, err)
	}
	s.gameChannels[id] = ch
	s.Unlock()

	if err := s.ps.Publish(s.gameUpdateChannel, &pubsub.Payload{
		Type: pubsub.PayloadTypeGameUpdated,
		GameUpdated: &pubsub.GameUpdated{
			ID:     id,
			Name:   name,
			Status: banana.WaitingForPlayers,
		},
	}); err != nil {
		log.Printf("failed to publish new game %q: %v", name, err)
	}

	return &pb.NewGameResponse{Id: string(id)}, nil
}

func (s *Server) ListGames(ctx context.Context, req *pb.ListGamesRequest) (*pb.ListGamesResponse, error) {
	gs, err := s.db.Games()
	if err != nil {
		return nil, fmt.Errorf("failed to get games: %v", err)
	}
	sort.Slice(gs, func(i, j int) bool { return gs[i].CreatedAt.Before(gs[j].CreatedAt) })

	var pbgs []*pb.Game
	for _, g := range gs {
		ps, err := s.db.Players(g.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get players: %v", err)
		}

		pbgs = append(pbgs, &pb.Game{
			Id:          string(g.ID),
			Name:        g.Name,
			Status:      gameStatusMap[g.Status],
			PlayerCount: int32(len(ps)),
		})
	}

	return &pb.ListGamesResponse{Games: pbgs}, nil
}

func (s *Server) StreamGames(req *pb.ListGamesRequest, stream pb.BananaService_StreamGamesServer) error {
	gs, err := s.db.Games()
	if err != nil {
		return fmt.Errorf("failed to get games: %v", err)
	}

	var pbgs []*pb.Game
	for _, g := range gs {
		ps, err := s.db.Players(g.ID)
		if err != nil {
			return fmt.Errorf("failed to get players: %v", err)
		}

		pbgs = append(pbgs, &pb.Game{
			Id:          string(g.ID),
			Name:        g.Name,
			Status:      gameStatusMap[g.Status],
			PlayerCount: int32(len(ps)),
		})
	}

	if err := stream.Send(&pb.GamesList{Games: pbgs}); err != nil {
		log.Printf("failed to send initial games list: %v", err)
	}

	sub, err := s.ps.Subscribe(s.gameUpdateChannel)
	if err != nil {
		log.Printf("failed to subscribe to new game updates: %v", err)
		return nil
	}
	defer sub.Close()

	for {
		msg, done := sub.Next()
		if done {
			return nil
		}
		if msg.Channel != s.gameUpdateChannel || msg.Payload.Type != pubsub.PayloadTypeGameUpdated {
			log.Printf("unexpected message from pubsub system: %#v", msg)
			continue
		}
		update := msg.Payload.GameUpdated
		stream.Send(&pb.GamesList{
			Games: []*pb.Game{
				&pb.Game{
					Id:          string(update.ID),
					Name:        update.Name,
					Status:      gameStatusMap[update.Status],
					PlayerCount: int32(update.PlayerCount),
				},
			},
		})
	}

	return nil
}

func (s *Server) StreamLogs(req *pb.StreamLogsRequest, stream pb.BananaService_StreamLogsServer) error {
	gID := banana.GameID(req.GameId)
	pID := banana.PlayerID(req.PlayerId)

	s.RLock()
	gch := s.gameChannels[gID]
	pch := s.playerChannels[pID]
	s.RUnlock()

	sub, err := s.ps.Subscribe(gch, pch)
	if err != nil {
		log.Printf("game ID: %q, player ID: %q", gID, pID)
		log.Printf("failed to subscribe to log updates: %v", err)
		return nil
	}
	defer sub.Close()

	for {
		msg, done := sub.Next()
		if done {
			return nil
		}

		var le *pb.LogEntry

		switch msg.Payload.Type {
		case pubsub.PayloadTypePlayerMove:
			le = s.playerMoveLogEntry(msg.Payload.PlayerMove)
		case pubsub.PayloadTypePlayerDump:
			le = s.playerDumpLogEntry(msg.Payload.PlayerDump)
		case pubsub.PayloadTypeSelfTileUpdate:
			le = s.playerPeelLogEntry(msg.Payload.SelfTileUpdate)
		}

		if le == nil {
			continue
		}

		if err := stream.Send(le); err != nil {
			log.Printf("log stream closed for player %q: %v", pID, err)
		}
	}

	return nil
}

func (s *Server) playerMoveLogEntry(pm *pubsub.PlayerMove) *pb.LogEntry {
	return &pb.LogEntry{
		Event: &pb.LogEntry_PlayerMove{
			PlayerMove: &pb.PlayerMove{
				PlayerId:       string(pm.ID),
				PlayerName:     pm.Name,
				Word:           pm.Word,
				WordValid:      pm.WordValid,
				BoardConnected: pm.BoardValid,
			},
		},
	}
}

func (s *Server) playerDumpLogEntry(pd *pubsub.PlayerDump) *pb.LogEntry {
	return &pb.LogEntry{
		Event: &pb.LogEntry_PlayerDump{
			PlayerDump: &pb.PlayerDump{
				PlayerId:   string(pd.ID),
				PlayerName: pd.Name,
			},
		},
	}
}

func (s *Server) playerPeelLogEntry(stu *pubsub.SelfTileUpdate) *pb.LogEntry {
	if stu.PeelFrom == "" {
		return nil
	}

	p, err := s.db.Player(stu.PeelFrom)
	if err != nil {
		return nil
	}

	return &pb.LogEntry{
		Event: &pb.LogEntry_PlayerPeel{
			PlayerPeel: &pb.PlayerPeel{
				PlayerId:   string(p.ID),
				PlayerName: p.Name,
			},
		},
	}
}

func (s *Server) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	gID := banana.GameID(req.Id)
	g, err := s.db.Game(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	if g.Status != banana.WaitingForPlayers {
		// Can't start an in-progess or finished game.
		return nil, fmt.Errorf("can't start game in state %q", g.Status)
	}

	ps, err := s.db.Players(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive players for game id %q: %v", req.Id, err)
	}

	if len(ps) < 2 {
		return nil, errors.New("can't start game with less than two people")
	}

	bunch, err := s.db.Bunch(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive bunch for game id %q: %v", req.Id, err)
	}

	numTiles := banana.StartingTileCount(len(ps), scaleFactor)

	var (
		otus    []*pubsub.OtherTileUpdate
		players = make(map[banana.PlayerID]*banana.Tiles)
	)

	for _, p := range ps {
		tls, err := bunch.RemoveN(numTiles, s.r)
		if err != nil {
			return nil, fmt.Errorf("failed to get %d tiles for player (%q, %q): %v", numTiles, p.ID, p.Name, err)
		}
		players[p.ID] = tls

		s.sendPlayerUpdate(p.ID, &pubsub.Payload{
			Type:           pubsub.PayloadTypeSelfTileUpdate,
			SelfTileUpdate: &pubsub.SelfTileUpdate{Tiles: tls},
		})
		otus = append(otus, &pubsub.OtherTileUpdate{
			ID:           p.ID,
			TilesInHand:  numTiles,
			TilesInBoard: 0, // Because the game is just starting
		})
	}

	if err := s.db.StartGame(gID, players, bunch); err != nil {
		return nil, err
	}

	s.sendGameUpdate(gID, &pubsub.Payload{
		Type:        pubsub.PayloadTypeGameStarted,
		GameStarted: &pubsub.GameStarted{},
	})

	s.sendGameUpdate(gID, &pubsub.Payload{
		Type: pubsub.PayloadTypeOtherTileUpdates,
		OtherTileUpdates: &pubsub.OtherTileUpdates{
			Updates:        otus,
			RemainingTiles: bunch.Count(),
		},
	})

	return &pb.StartGameResponse{}, nil
}

func (s *Server) JoinGame(req *pb.JoinGameRequest, stream pb.BananaService_JoinGameServer) error {
	gID := banana.GameID(req.Id)
	g, err := s.db.Game(gID)
	if err != nil {
		return fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	var (
		pID   banana.PlayerID
		pName string
	)

	players, err := s.db.Players(gID)
	if err != nil {
		return fmt.Errorf("failed to retreive players for game id %q: %v", gID, err)
	}

	// Try to add a new player if they weren't already in the game.
	if req.PlayerId == "" {
		if g.Status != banana.WaitingForPlayers {
			// Can't join an in-progess or finished game.
			return fmt.Errorf("can't join game in state %q", g.Status)
		}

		name := strings.TrimSpace(req.Name)
		if name == "" {
			return errors.New("must specify a player name")
		}
		pName = name

		if pID, err = s.db.AddPlayer(gID, name); err != nil {
			return err
		}
		if err := s.ps.Publish(s.gameUpdateChannel, &pubsub.Payload{
			Type: pubsub.PayloadTypeGameUpdated,
			GameUpdated: &pubsub.GameUpdated{
				ID:          g.ID,
				Name:        g.Name,
				Status:      g.Status,
				PlayerCount: len(players) + 1,
			},
		}); err != nil {
			log.Printf("failed to publish new game %q: %v", name, err)
		}

		s.sendGameUpdate(gID, &pubsub.Payload{
			Type: pubsub.PayloadTypePlayerJoined,
			PlayerJoined: &pubsub.PlayerJoined{
				ID:   pID,
				Name: name,
			},
		})

		s.Lock()
		pch, err := s.ps.NewChannel(string(pID))
		if err != nil {
			log.Printf("failed to create channel for player %q: %v", pID, err)
			s.Unlock()
			return fmt.Errorf("failed to create channel for player %q: %v", pName, err)
		}
		s.playerChannels[pID] = pch
		s.Unlock()
	} else {
		// Find the player in the game.
		found := false
		for _, p := range players {
			if p.ID == banana.PlayerID(req.PlayerId) {
				pID = p.ID
				pName = p.Name
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("failed to add or find player: %+v", req)
		}
	}

	wps, err := s.wirePlayers(gID)
	if err != nil {
		return err
	}

	bunch, err := s.db.Bunch(gID)
	if err != nil {
		return err
	}

	board, err := s.db.Board(pID)
	if err != nil {
		return err
	}

	tiles, err := s.db.Tiles(pID)
	if err != nil {
		return err
	}

	stream.Send(&pb.GameUpdate{
		Update: &pb.GameUpdate_CurrentStatus{
			CurrentStatus: &pb.CurrentStatus{
				YourId:         string(pID),
				Players:        wps,
				RemainingTiles: int32(bunch.Count()),
				Board:          boardToWire(board),
				AllTiles:       &pb.Tiles{Letters: tiles.AsList()},
				Status:         gameStatusMap[g.Status],
			},
		},
	})

	s.RLock()
	gch := s.gameChannels[gID]
	pch := s.playerChannels[pID]
	s.RUnlock()

	sub, err := s.ps.Subscribe(gch, pch)
	if err != nil {
		log.Printf("failed to subscribe to new game updates: %v", err)
		return nil
	}
	defer sub.Close()

	for {
		msg, done := sub.Next()
		if done {
			return nil
		}

		gameOver := false
		var update *pb.GameUpdate
		switch msg.Payload.Type {
		case pubsub.PayloadTypeGameStarted:
			update = &pb.GameUpdate{
				Update: &pb.GameUpdate_GameStarted{
					GameStarted: &pb.GameStarted{},
				},
			}
		case pubsub.PayloadTypeGameEnded:
			update = &pb.GameUpdate{
				Update: &pb.GameUpdate_GameEnded{
					GameEnded: &pb.GameEnded{
						// TODO: Populate the final standings.
					},
				},
			}
			gameOver = true
		case pubsub.PayloadTypePlayerJoined:
			update = &pb.GameUpdate{
				Update: &pb.GameUpdate_PlayerUpdate{
					PlayerUpdate: &pb.PlayerUpdate{
						Player: &pb.Player{
							Id:   string(msg.Payload.PlayerJoined.ID),
							Name: msg.Payload.PlayerJoined.Name,
						},
					},
				},
			}
		case pubsub.PayloadTypeSelfTileUpdate:
			tu := msg.Payload.SelfTileUpdate
			update = &pb.GameUpdate{
				Update: &pb.GameUpdate_SelfTileUpdate{
					SelfTileUpdate: &pb.SelfTileUpdate{
						AllTiles:      &pb.Tiles{Letters: tu.Tiles.AsList()},
						FromOtherPeel: tu.PeelFrom != pID && tu.PeelFrom != "",
					},
				},
			}
		case pubsub.PayloadTypeOtherTileUpdates:
			tu := msg.Payload.OtherTileUpdates
			update = &pb.GameUpdate{
				Update: &pb.GameUpdate_OtherTileUpdate{
					OtherTileUpdate: s.toOtherTileUpdates(tu.Updates, tu.RemainingTiles),
				},
			}
		}

		if update != nil {
			if err := stream.Send(update); err != nil {
				log.Printf("stream closed for player %q: %v", pName, err)
				return nil
			}
		}

		if gameOver {
			break
		}
	}

	return nil
}

func (s *Server) toOtherTileUpdates(otus []*pubsub.OtherTileUpdate, remainingTiles int) *pb.OtherTileUpdates {
	var out []*pb.OtherTileUpdate
	for _, otu := range otus {
		out = append(out, &pb.OtherTileUpdate{
			PlayerId:     string(otu.ID),
			TilesInHand:  int32(otu.TilesInHand),
			TilesInBoard: int32(otu.TilesInBoard),
		})
	}

	return &pb.OtherTileUpdates{
		Updates:        out,
		RemainingTiles: int32(remainingTiles),
	}
}

func (s *Server) Spectate(req *pb.SpectateRequest, stream pb.BananaService_SpectateServer) error {
	gID := banana.GameID(req.Id)

	ps, err := s.db.Players(gID)
	if err != nil {
		return err
	}

	// Stream an initial board for each player.
	for _, p := range ps {
		board, err := s.db.Board(p.ID)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.SpectateUpdate{
			PlayerId:   string(p.ID),
			PlayerName: p.Name,
			Board:      boardToWire(board),
		})
		if err != nil {
			log.Printf("failed to stream to spectator: %v", err)
			return fmt.Errorf("failed to stream to spectator: %v", err)
		}
	}

	s.RLock()
	ch := s.gameChannels[gID]
	s.RUnlock()

	sub, err := s.ps.Subscribe(ch)
	if err != nil {
		log.Printf("failed to subscribe to updates for game %q: %v", gID, err)
		return fmt.Errorf("failed to subscribe to updates for game %q: %v", gID, err)
	}
	defer sub.Close()

	for {
		msg, done := sub.Next()
		if done {
			return nil
		}
		if msg.Channel != ch {
			log.Printf("unexpected message from pubsub system: %#v", msg)
			continue
		}

		if msg.Payload.Type != pubsub.PayloadTypePlayerMove {
			continue
		}

		mv := msg.Payload.PlayerMove

		p, err := s.db.Player(mv.ID)
		if err != nil {
			log.Printf("failed to load player: %v", err)
			continue
		}

		board, err := s.db.Board(mv.ID)
		if err != nil {
			log.Printf("failed to load board: %v", err)
			continue
		}

		err = stream.Send(&pb.SpectateUpdate{
			PlayerId:   string(p.ID),
			PlayerName: p.Name,
			Board:      boardToWire(board),
		})
		if err != nil {
			log.Printf("failed to stream to spectator: %v", err)
			return fmt.Errorf("failed to stream to spectator: %v", err)
		}
	}
}

func (s *Server) sendPlayerUpdate(pID banana.PlayerID, up *pubsub.Payload) {
	s.RLock()
	ch := s.playerChannels[pID]
	s.RUnlock()

	if err := s.ps.Publish(ch, up); err != nil {
		log.Printf("failed to publish player update: %v", err)
	}
}

func (s *Server) sendGameUpdate(id banana.GameID, up *pubsub.Payload) {
	s.RLock()
	ch := s.gameChannels[id]
	s.RUnlock()

	if err := s.ps.Publish(ch, up); err != nil {
		log.Printf("failed to publish game update: %v", err)
	}
}

func (s *Server) UpdateBoard(ctx context.Context, req *pb.UpdateBoardRequest) (*pb.UpdateBoardResponse, error) {
	// Convert the board to our domain format.
	b := boardFromWire(req.Board)

	pID := banana.PlayerID(req.PlayerId)
	p, err := s.db.Player(pID)
	if err != nil {
		log.Printf("failed to get player %q: %v", req.PlayerId, err)
		return nil, fmt.Errorf("failed to get player %q: %v", req.PlayerId, err)
	}

	board, err := s.db.Board(pID)
	if err != nil {
		return nil, err
	}

	tiles, err := s.db.Tiles(pID)
	if err != nil {
		return nil, err
	}

	// Add existing board to tiles.
	bts, err := board.AsTiles()
	if err != nil {
		log.Printf("failed to get board tiles")
		return nil, fmt.Errorf("failed to get board tiles")
	}
	tiles.Add(bts)

	// Check if the board they sent is valid.
	bv, err := b.Validate(tiles, s.dict)
	if err != nil {
		log.Printf("failed to validate board: %v", err)
		return nil, fmt.Errorf("failed to validate board: %v", err)
	}

	// We'll write the board as long as they aren't using letters they don't
	// have.
	if len(bv.ExtraLetters) > 0 {
		log.Printf("used letters you don't have: %v", bv.ExtraLetters)
		return nil, fmt.Errorf("used letters you don't have: %v", bv.ExtraLetters)
	}

	if err := s.db.UpdateBoard(pID, b); err != nil {
		log.Printf("failed to update board %q: %v", pID, err)
		return nil, fmt.Errorf("failed to update board %q: %v", pID, err)
	}

	tilesInHand := b.Diff(tiles)
	if err := s.db.UpdateTiles(pID, tilesInHand); err != nil {
		log.Printf("failed to update tiles %q: %v", pID, err)
		return nil, fmt.Errorf("failed to update tiles %q: %v", pID, err)
	}

	// A board is peelable if all the words are valid
	peelable := len(bv.InvalidWords) == 0 && len(bv.UnusedLetters) == 0 && !bv.DetachedBoard

	gID := banana.GameID(req.Id)
	bunch, err := s.db.Bunch(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive bunch for game %q: %v", gID, err)
	}

	ps, err := s.db.Players(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive players for game %q: %v", gID, err)
	}

	pTiles := make(map[banana.PlayerID]*banana.Tiles)
	for _, p := range ps {
		tiles, err := s.db.Tiles(p.ID)
		if err != nil {
			return nil, err
		}
		pTiles[p.ID] = tiles
	}

	// If the board was valid, clear out their tiles and let everyone know
	// what's up.
	if peelable {
		var gameOver bool
		// If there are less tiles than players, end the game.
		if bunch.Count() < len(ps) {
			gameOver = true
		}

		if !gameOver {
			for _, p := range ps {
				newTile, err := bunch.RemoveN(1, s.r)
				if err != nil {
					return nil, err
				}
				pTiles[p.ID].Add(newTile)

				if err := s.db.UpdateTiles(p.ID, pTiles[p.ID]); err != nil {
					return nil, fmt.Errorf("failed to update player %q tiles: %v", p.ID, err)
				}

				// Send everyone their new tile set.
				s.sendPlayerUpdate(p.ID, &pubsub.Payload{
					Type: pubsub.PayloadTypeSelfTileUpdate,
					SelfTileUpdate: &pubsub.SelfTileUpdate{
						Tiles:    pTiles[p.ID],
						PeelFrom: pID,
					},
				})
			}
		}

		if err := s.db.UpdateBunch(gID, bunch); err != nil {
			return nil, fmt.Errorf("UpdateBunch: %v", err)
		}

		if gameOver {
			if err := s.endGame(gID, p.Name); err != nil {
				return nil, fmt.Errorf("endGame: %v", err)
			}
		}
	}

	var (
		latestWord string
		validWord  = true
	)
	if req.LatestWord != nil && utf8.RuneCountInString(req.LatestWord.Text) > 1 {
		latestWord = req.LatestWord.Text
		for _, iw := range bv.InvalidWords {
			if iw.Word == latestWord {
				validWord = false
				break
			}
		}
	}

	// Send everyone what this player played.
	s.sendGameUpdate(gID, &pubsub.Payload{
		Type: pubsub.PayloadTypePlayerMove,
		PlayerMove: &pubsub.PlayerMove{
			ID:           p.ID,
			Name:         p.Name,
			Word:         latestWord,
			WordValid:    validWord,
			BoardValid:   !bv.DetachedBoard,
			TilesInBunch: bunch.Count(),
		},
	})

	var otus []*pubsub.OtherTileUpdate
	for _, p := range ps {
		board, err := s.db.Board(p.ID)
		if err != nil {
			return nil, err
		}

		bc, err := board.Count()
		if err != nil {
			return nil, err
		}

		otus = append(otus, &pubsub.OtherTileUpdate{
			ID:           p.ID,
			TilesInHand:  pTiles[p.ID].Count(),
			TilesInBoard: bc,
		})
	}

	// Send everyone's tile counts down.
	s.sendGameUpdate(gID, &pubsub.Payload{
		Type: pubsub.PayloadTypeOtherTileUpdates,
		OtherTileUpdates: &pubsub.OtherTileUpdates{
			Updates:        otus,
			RemainingTiles: bunch.Count(),
		},
	})

	// Convert the status to the wire format.
	return &pb.UpdateBoardResponse{
		InvalidWords:  charLocsListToWire(bv.InvalidWords),
		UnusedLetters: bv.UnusedLetters,
		DetachedBoard: bv.DetachedBoard,
	}, nil
}

func (s *Server) endGame(gID banana.GameID, name string) error {
	s.sendGameUpdate(gID, &pubsub.Payload{
		Type:      pubsub.PayloadTypeGameEnded,
		GameEnded: &pubsub.GameEnded{Winner: name},
	})

	return s.db.EndGame(gID)
}

// Dump exchanges a player's letter for three from the bunch
func (s *Server) Dump(ctx context.Context, req *pb.DumpRequest) (*pb.DumpResponse, error) {
	gID := banana.GameID(req.Id)
	pID := banana.PlayerID(req.PlayerId)

	p, err := s.db.Player(pID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive player id %q: %v", req.PlayerId, err)
	}

	g, err := s.db.Game(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	if g.Status != banana.InProgress {
		// Can't dump in a game that isn't ongoing.
		return nil, fmt.Errorf("can't dump in a game in state %q", g.Status)
	}

	if len(req.Letter) != 1 {
		return nil, errors.New("can't dump more than one letter at a time")
	}

	letter := banana.Letter(req.Letter[0])

	bunch, err := s.db.Bunch(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive bunch: %v", err)
	}

	if bunch.Count() < DumpSize {
		// We don't have enough tiles to give them
		return nil, errors.New("not enough tiles to dump")
	}

	board, err := s.db.Board(pID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive board: %v", err)
	}

	tiles, err := s.db.Tiles(pID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive tiles: %v", err)
	}

	if tiles.Freq(letter) <= 0 {
		// They don't actually even have this letter to give away.
		return nil, fmt.Errorf("you don't even have %q in your hand", letter)
	}

	// If we're here, it's probably safe to go ahead with the dump.

	// Take the shit letter from them.
	tiles.Dec(letter)

	// Give them three tiles, before we throw their shit tile back in.
	tls, err := bunch.RemoveN(DumpSize, s.r)
	if err != nil {
		return nil, err
	}
	tiles.Add(tls)

	// We put their letter back in the pot after we take out their shit letter.
	bunch.Inc(letter)

	if err := s.db.UpdateTiles(pID, tiles); err != nil {
		return nil, err
	}

	if err := s.db.UpdateBunch(gID, bunch); err != nil {
		return nil, err
	}

	s.sendPlayerUpdate(pID, &pubsub.Payload{
		Type:           pubsub.PayloadTypeSelfTileUpdate,
		SelfTileUpdate: &pubsub.SelfTileUpdate{Tiles: tiles},
	})

	bc, err := board.Count()
	if err != nil {
		return nil, err
	}

	s.sendGameUpdate(gID, &pubsub.Payload{
		Type: pubsub.PayloadTypeOtherTileUpdates,
		OtherTileUpdates: &pubsub.OtherTileUpdates{
			Updates: []*pubsub.OtherTileUpdate{
				&pubsub.OtherTileUpdate{
					ID:           pID,
					TilesInBoard: bc,
					TilesInHand:  tiles.Count(),
				},
			},
		},
	})

	s.sendGameUpdate(gID, &pubsub.Payload{
		Type: pubsub.PayloadTypePlayerDump,
		PlayerDump: &pubsub.PlayerDump{
			ID:   pID,
			Name: p.Name,
		},
	})

	return &pb.DumpResponse{
		AllTiles: &pb.Tiles{
			Letters: tiles.AsList(),
		},
	}, nil
}

func (s *Server) wirePlayers(gID banana.GameID) ([]*pb.Player, error) {
	players, err := s.db.Players(gID)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive players for game id %q: %v", gID, err)
	}

	var out []*pb.Player
	for _, p := range players {
		board, err := s.db.Board(p.ID)
		if err != nil {
			return nil, err
		}

		tob, err := board.Count()
		if err != nil {
			return nil, err
		}

		tiles, err := s.db.Tiles(p.ID)
		if err != nil {
			return nil, err
		}

		out = append(out, &pb.Player{
			Id:           string(p.ID),
			Name:         p.Name,
			TilesInHand:  int32(tiles.Count()),
			TilesInBoard: int32(tob),
		})
	}

	return out, nil
}
