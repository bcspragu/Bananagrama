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

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

const (
	DumpSize    = 3
	scaleFactor = 1
	// TODO: Add MaxPlayers back in
)

type Server struct {
	r    *rand.Rand
	db   banana.DB
	dict banana.Dictionary

	sync.RWMutex
	updates map[banana.GameID]map[banana.PlayerID]chan *pb.GameUpdate
}

func New(r *rand.Rand, db banana.DB, dict banana.Dictionary) *Server {
	return &Server{
		r:       r,
		db:      db,
		dict:    dict,
		updates: make(map[banana.GameID]map[banana.PlayerID]chan *pb.GameUpdate),
	}
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

	id, err := s.db.NewGame(req.Name, bunch)
	if err != nil {
		return nil, err
	}
	s.Lock()
	s.updates[id] = make(map[banana.PlayerID]chan *pb.GameUpdate)
	s.Unlock()
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
		pbgs = append(pbgs, &pb.Game{
			Id:          string(g.ID),
			Name:        g.Name,
			Status:      gameStatusMap[g.Status],
			PlayerCount: int32(len(g.Players)),
		})
	}

	return &pb.ListGamesResponse{Games: pbgs}, nil
}

func (s *Server) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	gid := banana.GameID(req.Id)
	g, err := s.db.Game(gid)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	if g.Status != banana.WaitingForPlayers {
		// Can't start an in-progess or finished game.
		return nil, fmt.Errorf("can't start game in state %q", g.Status)
	}

	numTiles := banana.StartingTileCount(len(g.Players), scaleFactor)
	players := make(map[banana.PlayerID]*banana.Tiles)
	for _, p := range g.Players {
		tls, err := g.Bunch.RemoveN(numTiles, s.r)
		if err != nil {
			return nil, fmt.Errorf("failed to get %d tiles for player (%q, %q): %v", numTiles, p.ID, p.Name, err)
		}
		players[p.ID] = tls
	}

	if err := s.db.StartGame(gid, players, g.Bunch); err != nil {
		return nil, err
	}

	s.updateForGame(gid, &pb.GameUpdate{
		Update: &pb.GameUpdate_StatusUpdate{
			StatusUpdate: &pb.StatusUpdate{Status: pb.StatusUpdate_GAME_STARTED},
		},
	})

	if err := s.sendTiles(gid, &pb.TileUpdate{Event: pb.TileUpdate_SPLIT}); err != nil {
		return nil, fmt.Errorf("failed to send tiles: %v", err)
	}

	if err := s.sendPlayers(gid); err != nil {
		return nil, fmt.Errorf("failed to send players: %v", err)
	}

	return &pb.StartGameResponse{}, nil
}

func (s *Server) JoinGame(req *pb.JoinGameRequest, stream pb.BananaService_JoinGameServer) error {
	gid := banana.GameID(req.Id)
	g, err := s.db.Game(gid)
	if err != nil {
		return fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	var (
		pChan chan *pb.GameUpdate
		pid   banana.PlayerID
		pName string
	)

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

		if pid, err = s.db.AddPlayer(gid, name); err != nil {
			return err
		}
		s.Lock()
		if _, ok := s.updates[gid]; !ok {
			s.Unlock()
			return fmt.Errorf("game %q not found", gid)
		}
		pChan = make(chan *pb.GameUpdate)
		s.updates[gid][pid] = pChan
		s.Unlock()
		log.Printf("Added player %q to game %q", pid, req.Id)
	} else {
		// Find the player in the game.
		for _, p := range g.Players {
			if p.ID == banana.PlayerID(req.PlayerId) {
				// TODO: Maybe add a check that updates[gid] exists, like above.
				s.RLock()
				pChan = s.updates[gid][p.ID]
				pid = p.ID
				pName = p.Name
				s.RUnlock()
				break
			}
		}
		log.Printf("Player %q rejoined game %q", pid, req.Id)

		// TODO: Send them the info they need to get back up to speed.
	}

	if pChan == nil {
		return fmt.Errorf("failed to add or find player: %+v", req)
	}

	stream.Send(&pb.GameUpdate{
		Update: &pb.GameUpdate_YouUpdate{
			YouUpdate: &pb.YouUpdate{YourId: string(pid)},
		},
	})

	p, err := s.db.Player(pid)
	if err != nil {
		return err
	}

	stream.Send(&pb.GameUpdate{
		Update: &pb.GameUpdate_BoardUpdate{
			BoardUpdate: &pb.BoardUpdate{Board: boardToWire(p.Board)},
		},
	})

	stream.Send(&pb.GameUpdate{
		Update: &pb.GameUpdate_TileUpdate{
			TileUpdate: &pb.TileUpdate{
				Event:    pb.TileUpdate_JOIN,
				Player:   p.Name,
				AllTiles: &pb.Tiles{Letters: p.Tiles.AsList()},
			},
		},
	})

	go func() {
		s.updateForGame(gid, &pb.GameUpdate{
			Update: &pb.GameUpdate_StatusUpdate{
				StatusUpdate: &pb.StatusUpdate{Status: pb.StatusUpdate_GAME_STARTED},
			},
		})

		if err := s.sendPlayers(gid); err != nil {
			log.Printf("Failed to send player list for game %q: %v", gid, err)
		}
	}()

	for {
		update := <-s.updates[gid][pid]

		gameOver := false
		switch u := update.Update.(type) {
		case *pb.GameUpdate_PlayerUpdate:
			// Handle PlayerUpdate.
		case *pb.GameUpdate_StatusUpdate:
			su := u.StatusUpdate
			// Handle StatusUpdate.
			if su.Status == pb.StatusUpdate_GAME_OVER {
				gameOver = true
			}
		case *pb.GameUpdate_TileUpdate:
			// Handle TileUpdate.
		}

		if err := stream.Send(update); err != nil {
			log.Printf("stream send error: %v", err)
			return nil
		}
		log.Printf("successfully sent to %q: %#v", pName, update)

		if gameOver {
			break
		}
	}

	return nil
}

func (s *Server) sendMove(id banana.GameID, name, word string) {
	s.updateForGame(id, &pb.GameUpdate{
		Update: &pb.GameUpdate_MoveUpdate{
			MoveUpdate: &pb.MoveUpdate{
				Player: name,
				Word:   word,
			},
		},
	})
}

func (s *Server) sendPlayers(id banana.GameID) error {
	g, err := s.db.Game(id)
	if err != nil {
		return err
	}

	var players []*pb.Player
	for _, p := range g.Players {
		bc, err := p.Board.Count()
		if err != nil {
			return fmt.Errorf("failed to get board count: %v", err)
		}

		players = append(players, &pb.Player{
			Name:         p.Name,
			TilesInHand:  int32(p.Tiles.Count()),
			TilesInBunch: int32(bc),
		})
		fmt.Printf("%+v\n", players[len(players)-1])
	}

	up := &pb.GameUpdate{
		Update: &pb.GameUpdate_PlayerUpdate{
			PlayerUpdate: &pb.PlayerUpdate{
				Players:        players,
				RemainingTiles: int32(g.Bunch.Count()),
			},
		},
	}

	s.updateForGame(id, up)
	return nil
}

func (s *Server) sendTiles(id banana.GameID, base *pb.TileUpdate) error {
	g, err := s.db.Game(id)
	if err != nil {
		return err
	}

	s.RLock()
	for _, p := range g.Players {
		s.updates[id][p.ID] <- &pb.GameUpdate{
			Update: &pb.GameUpdate_TileUpdate{
				TileUpdate: &pb.TileUpdate{
					Event:    base.Event,
					Player:   base.Player,
					AllTiles: &pb.Tiles{Letters: p.Tiles.AsList()},
				},
			},
		}
	}
	s.RUnlock()
	return nil
}

func (s *Server) updateForGame(id banana.GameID, up *pb.GameUpdate) {
	s.RLock()
	for _, pc := range s.updates[id] {
		pc <- up
	}
	s.RUnlock()
}

func (s *Server) UpdateBoard(ctx context.Context, req *pb.UpdateBoardRequest) (*pb.UpdateBoardResponse, error) {
	gid := banana.GameID(req.Id)
	// Convert the board to our domain format.
	b := boardFromWire(req.Board)

	pid := banana.PlayerID(req.PlayerId)
	p, err := s.db.Player(pid)
	if err != nil {
		log.Printf("failed to get player %q: %v", req.PlayerId, err)
		return nil, fmt.Errorf("failed to get player %q: %v", req.PlayerId, err)
	}

	// Add existing board to tiles.
	tiles := p.Tiles.Clone()
	bts, err := p.Board.AsTiles()
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

	if err := s.db.UpdatePlayer(pid, b, b.Diff(tiles)); err != nil {
		log.Printf("failed to update player %q: %v", pid, err)
		return nil, fmt.Errorf("failed to update player %q: %v", pid, err)
	}
	log.Printf("Player %q updated their board, result %+v", p.Name, bv)

	// A board is peelable if all the words are valid
	peelable := len(bv.InvalidWords) == 0 && len(bv.UnusedLetters) == 0 && !bv.DetachedBoard

	// If the board was valid, clear out their tiles and let everyone know what's
	// up.
	if peelable {
		log.Printf("Peel by %q", p.Name)
		if err := s.issuePeel(gid, p); err != nil {
			log.Printf("failed to issue peels: %v", err)
			return nil, fmt.Errorf("failed to issue peels: %v", err)
		}
	}

	if err := s.sendPlayers(gid); err != nil {
		log.Printf("failed to update players: %v", err)
		return nil, fmt.Errorf("failed to update players: %v", err)
	}

	if req.LatestWord != nil {
		s.sendMove(gid, p.Name, req.LatestWord.Text)
	}

	// Convert the status to the wire format.
	return &pb.UpdateBoardResponse{
		InvalidWords:  charLocsListToWire(bv.InvalidWords),
		UnusedLetters: bv.UnusedLetters,
		DetachedBoard: bv.DetachedBoard,
	}, nil
}

func (s *Server) issuePeel(id banana.GameID, p *banana.Player) error {
	s.RLock()
	defer s.RUnlock()

	g, err := s.db.Game(id)
	if err != nil {
		return fmt.Errorf("failed to retreive game %q: %v", id, err)
	}

	// If there are less tiles than players, end the game.
	if g.Bunch.Count() < len(g.Players) {
		return s.endGame(id)
	}

	for pid, c := range s.updates[id] {
		sp, err := s.db.Player(pid)
		if err != nil {
			return fmt.Errorf("failed to retreive player %q: %v", pid, err)
		}

		tiles, err := g.Bunch.RemoveN(1, s.r)
		if err != nil {
			return err
		}
		sp.Tiles.Add(tiles)

		if err := s.db.UpdatePlayer(pid, sp.Board, sp.Tiles); err != nil {
			return fmt.Errorf("failed to update player %q board: %v", pid, err)
		}

		log.Printf("Sending tiles to %q: %+v", sp.Name, sp.Tiles.AsList())
		c <- &pb.GameUpdate{
			Update: &pb.GameUpdate_TileUpdate{
				TileUpdate: &pb.TileUpdate{
					Event:    pb.TileUpdate_PEEL,
					Player:   p.Name,
					AllTiles: &pb.Tiles{Letters: sp.Tiles.AsList()},
				},
			},
		}
	}

	if err := s.db.UpdateBunch(id, g.Bunch); err != nil {
		return fmt.Errorf("UpdateBunch: %v", err)
	}

	return nil
}

func (s *Server) endGame(id banana.GameID) error {
	for _, c := range s.updates[id] {
		c <- &pb.GameUpdate{
			Update: &pb.GameUpdate_StatusUpdate{
				StatusUpdate: &pb.StatusUpdate{Status: pb.StatusUpdate_GAME_OVER},
			},
		}
	}
	return s.db.EndGame(id)
}

// Dump exchanges a player's letter for three from the bunch
func (s *Server) Dump(ctx context.Context, req *pb.DumpRequest) (*pb.DumpResponse, error) {
	gid := banana.GameID(req.Id)
	pid := banana.PlayerID(req.PlayerId)
	g, err := s.db.Game(gid)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	p, err := s.db.Player(pid)
	if err != nil {
		return nil, fmt.Errorf("failed to retreive player id %q: %v", req.PlayerId, err)
	}

	if g.Status != banana.InProgress {
		// Can't dump in a game that isn't ongoing.
		return nil, fmt.Errorf("can't dump in a game in state %q", g.Status)
	}

	if len(req.Letter) != 1 {
		return nil, errors.New("can't dump more than one letter at a time")
	}

	letter := banana.Letter(req.Letter[0])

	if g.Bunch.Count() < DumpSize {
		// We don't have enough tiles to give them
		return nil, errors.New("not enough tiles to dump")
	}

	if p.Tiles.Freq(letter) <= 0 {
		// They don't actually even have this letter to give away.
		return nil, fmt.Errorf("you don't even have %q in your hand", letter)
	}

	// If we're here, it's probably safe to go ahead with the dump.

	// Take the shit letter from them.
	p.Tiles.Dec(letter)

	// Give them three tiles, before we throw their shit tile back in.
	tls, err := g.Bunch.RemoveN(DumpSize, s.r)
	if err != nil {
		return nil, err
	}
	p.Tiles.Add(tls)

	// We put their letter back in the pot after we take out their shit letter.
	g.Bunch.Inc(letter)

	if err := s.db.UpdatePlayer(pid, p.Board, p.Tiles); err != nil {
		return nil, err
	}

	if err := s.db.UpdateBunch(gid, g.Bunch); err != nil {
		return nil, err
	}

	s.RLock()
	s.updates[gid][pid] <- &pb.GameUpdate{
		Update: &pb.GameUpdate_TileUpdate{
			TileUpdate: &pb.TileUpdate{
				Event:    pb.TileUpdate_DUMP,
				Player:   p.Name,
				AllTiles: &pb.Tiles{Letters: p.Tiles.AsList()},
			},
		},
	}
	s.RUnlock()

	if err := s.sendPlayers(gid); err != nil {
		return nil, err
	}

	log.Printf("Player %q dumped a %q", p.Name, letter)

	return &pb.DumpResponse{}, nil
}
