package main

import (
	"fmt"
	"sync"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
	capnp "zombiezen.com/go/capnproto2"
)

var (
	wireOrientationMap = map[potassium.Word_Orientation]engine.Orientation{
		potassium.Word_Orientation_unknown:    engine.None,
		potassium.Word_Orientation_horizontal: engine.Horizontal,
		potassium.Word_Orientation_vertical:   engine.Vertical,
	}

	engineStatusMap = map[engine.BoardStatusCode]potassium.PeelResponse_Status{
		engine.Success:       potassium.PeelResponse_Status_success,
		engine.InvalidWord:   potassium.PeelResponse_Status_invalidWord,
		engine.DetachedBoard: potassium.PeelResponse_Status_detachedBoard,
		engine.NotAllLetters: potassium.PeelResponse_Status_notAllLetters,
		engine.ExtraLetters:  potassium.PeelResponse_Status_extraLetters,
	}
)

type player struct {
	potassium.Player
	name string
}

type game struct {
	e *aiEndpoint

	mu     sync.RWMutex
	score  int
	tiles  engine.FreqList
	player *player
}

func (g *game) Peel(call potassium.Game_peel) error {
	req := call.Params
	resp := call.Results

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
	g.mu.RLock()
	status := b.ValidateBoard(g.tiles)
	g.mu.RUnlock()

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
		case engine.NotAllLetters:
			err = resp.SetNotAllLetters(tl)
		case engine.ExtraLetters:
			err = resp.SetExtraLetters(tl)
		case engine.InvalidWord:
			err = resp.SetInvalidWord(tl)
		}

		if err != nil {
			return fmt.Errorf("receiving peel: error setting response errors: %v", err)
		}
	}

	// If the board was valid, increment the players score and let them know waz good
	if status.Code == engine.Success {
		// Send the peel data over for persistence/sending to clients
		g.e.peelChan <- &peelInfo{
			player: g.player.name,
			board:  wb,
		}
	}

	return nil
}

type peelInfo struct {
	player string
	board  potassium.Board
}

// Dump exchanges a player's letter for three from the bunch
func (g *game) Dump(call potassium.Game_dump) error {
	// TODO(bsprague): Check all of this locking stuff
	g.e.mu.Lock()
	defer g.e.mu.Unlock()

	req := call.Params
	resp := call.Results

	// The letter the player wants to dump
	l, err := req.Letter()
	if err != nil {
		return err
	}

	if len(l) != 1 {
		resp.SetStatus(potassium.DumpResponse_Status_malformedRequest)
		return nil
	}

	letter := engine.Letter(l[0])

	// We don't have enough tiles to give them
	if g.e.bunch.Count() < DumpSize {
		resp.SetStatus(potassium.DumpResponse_Status_notEnoughTiles)
		return nil
	}

	g.mu.RLock()
	if g.tiles.Freq(letter) <= 0 {
		// They don't actually even have this letter to give away
		resp.SetStatus(potassium.DumpResponse_Status_letterNotInTiles)
		g.mu.RUnlock()
		return nil
	}
	g.mu.RUnlock()

	// If we're here, it's probably safe to go ahead with the dump. We'll start
	// with a write lock
	g.mu.Lock()
	// Take the shit letter from them
	g.tiles.Dec(letter)
	// Give them three tiles, before we throw their shit tile back in
	letters := make([]engine.Letter, DumpSize)
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

func boardFromWire(b potassium.Board) (*engine.Board, error) {
	wireWords, err := b.Words()
	if err != nil {
		return nil, err
	}

	words := make([]engine.Word, wireWords.Len())
	for i := 0; i < wireWords.Len(); i++ {
		w, err := wordFromWire(wireWords.At(i))
		if err != nil {
			return nil, err
		}
		words[i] = w
	}

	return &engine.Board{
		Words:      words,
		Dictionary: dict,
	}, nil
}

func wordFromWire(w potassium.Word) (engine.Word, error) {
	text, err := w.Text()
	if err != nil {
		return engine.Word{}, nil
	}

	return engine.Word{
		Orientation: wireOrientationMap[w.Orientation()],
		Text:        text,
		Loc:         engine.Loc{X: int(w.X()), Y: int(w.Y())},
	}, nil
}
