package srv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

const (
	DumpSize = 3
)

type Server struct {
	r       *rand.Rand
	db      banana.DB
	dict    banana.Dictionary
	updates map[banana.PlayerID]chan *pb.GameUpdate
}

func New(r *rand.Rand, db banana.DB, dict banana.Dictionary) *Server {
	return &Server{
		r:       r,
		db:      db,
		dict:    dict,
		updates: make(map[banana.PlayerID]chan *pb.GameUpdate),
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
	return &pb.NewGameResponse{Id: string(id)}, nil
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

	numTiles := banana.StartingTileCount(len(g.Players), req.ScaleFactor)
	players := make(map[banana.PlayerID]*Tiles)
	for id, p := range g.Players {
		tls, err := g.Bunch.RemoveN(numTiles)
		if err != nil {
			return nil, fmt.Errorf("failed to get %d tiles for player (%q, %q): %v", numTiles, p.ID, p.Name, err)
		}
		players[id] = tls
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
	pid, err := s.db.AddPlayer(banana.GameID(req.Id), req.Name)
	if err != nil {
		return err
	}
	if _, ok := s.updates[pid]; !ok {
		s.updates[pid] = make(chan *pb.GameUpdate)
	}

	log.Println("Added player %q to game %q", pid, req.Id)

	for {
		update := <-s.updates[pid]

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
	// Load board player sent us
	wb, err := req.Board()
	if err != nil {
		return err
	}

	// Convert the board to a friendlier format
	b, err := boardFromWire(wb)
	if err != nil {
		return err
	}

	// Check if the board they sent is valid
	status := b.ValidateBoard(g.tiles)

	// Convert the status to the wire format
	resp.SetStatus(engineStatusMap[status.Code])
	// If our error code came with some errors, drop those bad boys into the response
	if len(status.Errors) > 0 {
		_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			return err
		}
		tl, err := capnp.NewTextList(seg, int32(len(status.Errors)))
		if err != nil {
			return err
		}

		for i, wordErrs := range status.Errors {
			err = tl.Set(i, wordErrs)
			if err != nil {
				return fmt.Errorf("receiving peel: error setting errors %v: %v", wordErrs, err)
			}
		}
		switch status.Code {
		case banana.NotAllLetters:
			err = resp.SetNotAllLetters(tl)
		case banana.ExtraLetters:
			err = resp.SetExtraLetters(tl)
		case banana.InvalidWord:
			err = resp.SetInvalidWord(tl)
		}

		if err != nil {
			return fmt.Errorf("receiving peel: error setting response errors: %v", err)
		}
	}

	// If the board was valid, increment the players score and let them know waz good
	if status.Code == banana.Success {
		// Send the peel data over for persistence/sending to clients
		g.e.peelChan <- &peelInfo{
			player: g.player.name,
			board:  wb,
		}
	}

	return nil
}

// Dump exchanges a player's letter for three from the bunch
func (s *Server) Dump(ctx context.Context, req *pb.DumpRequest) (*pb.DumpResponse, error) {
	g, err := s.db.Game(banana.GameID(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to retreive game id %q: %v", req.Id, err)
	}

	if g.Status != banana.InProgress {
		// Can't dump in a game that isn't ongoing.
		return nil, fmt.Errorf("can't dump in a game in state %q", g.Status)
	}

	if len(req.Letter) != 1 {
		return nil, fmt.Errorf("can't dump more than one letter at a time", g.Status)
	}

	letter := banana.Letter(l[0])

	if g.bunch.Count() < DumpSize {
		// We don't have enough tiles to give them
		return nil, errors.New("not enough tiles to dump")
	}

	if g.Bunch.Freq(letter) <= 0 {
		// They don't actually even have this letter to give away.
		return nil, fmt.Errorf("you don't even have %q in your hand", letter)
	}

	// If we're here, it's probably safe to go ahead with the dump. We'll start
	// with a write lock
	// TODO: Fix everything below here, it's still all old code.

	// Take the shit letter from them
	g.tiles.Dec(letter)
	// Give them three tiles, before we throw their shit tile back in
	letters := make([]banana.Letter, DumpSize)
	for i := 0; i < DumpSize; i++ {
		letters[i] = g.e.bunch.Tile()
		g.tiles.Inc(letters[i])
	}
	// We put their letter back in the pot after we take out their shit letter
	g.e.bunch.Inc(letter)

	g.mu.Unlock()
	tl, err := resp.NewLetters(DumpSize)
	if err != nil {
		return err
	}
	for i, letter := range letters {
		err = tl.Set(i, letter.String())
		if err != nil {
			return fmt.Errorf("receiving dump: error setting new letters: %v", err)
		}
	}
	resp.SetStatus(potassium.DumpResponse_Status_success)

	// Save the dump to the DB
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return fmt.Errorf("saving peel: error making message: %v", err)
	}

	// Create the dump
	d, err := potassium.NewRootDump(seg)
	if err != nil {
		return fmt.Errorf("saving dump: error making message: %v", err)
	}

	err = d.SetPlayer(g.player.name)
	if err != nil {
		return fmt.Errorf("saving dump: error setting player name: %v", err)
	}

	err = d.SetDump(tl)
	if err != nil {
		return fmt.Errorf("saving dump: error setting dump tile list: %v", err)
	}

	go db.addDump(d)

	return nil
}
