package srv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

const (
	DumpSize = 3
	// TODO: Add MaxPlayers back in
)

type Server struct {
	r    *rand.Rand
	db   banana.DB
	dict banana.Dictionary

	*sync.RWMutex
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
	id, err := s.db.NewGame(req.Name)
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

	var pbgs []*pb.Game
	for id, g := range gs {
		pbgs = append(pbgs, &pb.Game{
			Id:   string(id),
			Name: g.Name,
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
		// Can't join an in-progess or finished game.
		return nil, fmt.Errorf("can't join game in state %q", g.Status)
	}

	b, err := banana.NewBunch(banana.Bananagrams(), 1)
	if err != nil {
		return nil, fmt.Errorf("failed to make bunch for game: %v", err)
	}

	numTiles := banana.StartingTileCount(len(g.Players), int(req.ScaleFactor))
	players := make(map[banana.PlayerID]*banana.Tiles)
	for _, p := range g.Players {
		tls, err := g.Bunch.RemoveN(numTiles, s.r)
		if err != nil {
			return nil, fmt.Errorf("failed to get %d tiles for player (%q, %q): %v", numTiles, p.ID, p.Name, err)
		}
		players[p.ID] = tls
	}

	if err := s.db.StartGame(gid, players, b); err != nil {
		return nil, err
	}
	return &pb.StartGameResponse{}, nil
}

func (s *Server) JoinGame(req *pb.JoinGameRequest, stream pb.BananaService_JoinGameServer) error {
	g, err := s.db.Game(banana.GameID(req.Id))
	if err != nil {
		return fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	if g.Status != banana.WaitingForPlayers {
		// Can't join an in-progess or finished game.
		return fmt.Errorf("can't join game in state %q", g.Status)
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return errors.New("must specify a player name")
	}
	gid := banana.GameID(req.Id)
	pid, err := s.db.AddPlayer(gid, req.Name)
	if err != nil {
		return err
	}
	s.Lock()
	pm, ok := s.updates[gid]
	if !ok {
		s.Unlock()
		return fmt.Errorf("game %q not found", gid)
	}
	pm[pid] = make(chan *pb.GameUpdate)
	s.updates[gid] = pm
	s.Unlock()

	log.Printf("Added player %q to game %q", pid, req.Id)

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

		stream.Send(update)

		if gameOver {
			break
		}
	}

	return nil
}

func (s *Server) Peel(ctx context.Context, req *pb.PeelRequest) (*pb.PeelResponse, error) {
	gid := banana.GameID(req.Id)
	// Convert the board to our domain format.
	b := boardFromWire(req.Board, s.dict)

	pid := banana.PlayerID(req.PlayerId)
	p, err := s.db.Player(pid)
	if err != nil {
		return nil, fmt.Errorf("failed to get player %q: %v", req.PlayerId, err)
	}

	// Check if the board they sent is valid
	status := b.ValidateBoard(p.Tiles)

	// If the board was valid, clear out their tiles and let everyone know what's
	// up.
	if status.Code == banana.Success {
		if err := s.db.UpdatePlayer(pid, b, banana.NewTiles()); err != nil {
			return nil, fmt.Errorf("failed to update player %q board: %v", pid, err)
		}
		if err := s.issuePeel(gid, p.Name); err != nil {
			return nil, fmt.Errorf("failed to issue peels: %v", err)
		}
	}

	// Convert the status to the wire format.
	return &pb.PeelResponse{
		Status: engineStatusMap[status.Code],
		Errors: status.Errors,
	}, nil
}

func (s *Server) issuePeel(id banana.GameID, name string) error {
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
		p, err := s.db.Player(pid)
		if err != nil {
			return fmt.Errorf("failed to retreive player %q: %v", pid, err)
		}

		tiles, err := g.Bunch.RemoveN(1, s.r)
		if err != nil {
			return err
		}
		p.Tiles.Add(tiles)

		if err := s.db.UpdatePlayer(pid, p.Board, p.Tiles); err != nil {
			return fmt.Errorf("failed to update player %q board: %v", pid, err)
		}

		c <- &pb.GameUpdate{
			Update: &pb.GameUpdate_TileUpdate{
				TileUpdate: &pb.TileUpdate{
					Event:    pb.TileUpdate_PEEL,
					Player:   name,
					AllTiles: &pb.Tiles{Letters: p.Tiles.AsList()},
				},
			},
		}
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

	return &pb.DumpResponse{}, nil
}
