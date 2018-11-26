package srv

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

type Server struct {
	db      banana.DB
	dict    banana.Dictionary
	updates map[banana.PlayerID]chan *pb.GameUpdate
}

func New(db banana.DB, dict banana.Dictionary) *Server {
	return &Server{
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

func (s *Server) JoinGame(req *pb.JoinGameRequest, stream pb.BananaService_JoinGameServer) error {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("must specify a player name")
	}
	pid, err := s.AddPlayer(banana.GameID(req.Id), req.Name)
	if err != nil {
		return nil, err
	}
	if _, ok := s.updates[pid]; !ok {
		s.updates[pid] = make(chan *pb.GameUpdate)
	}

	log.Println("Added player %q to game %q", pid, req.Id)

	for {
		update <- s.updates[pid]

		gameOver := false
		switch u := update.(type) {
		case *pb.PlayerUpdate:
			// Handle PlayerUpdate.
		case *pb.StatusUpdate:
			// Handle StatusUpdate.
			if u.Status == pb.StatusUpdate_GAME_OVER {
				gameOver = true
			}
		case *pb.TileUpdate:
			// Handle TileUpdate.
		}

		stream.Send(update)

		if gameOver {
			break
		}
	}
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

func (s *Server) Dump(ctx context.Context, req *pb.DumpRequest) (*pb.DumpResponse, error) {
	return nil, errors.New("not implemented")
}
