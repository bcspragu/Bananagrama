package main

import (
	"log"
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

func (g *game) AddTile(l engine.Letter) {
	g.mu.Lock()
	g.tiles.Inc(l)
	g.mu.Unlock()
}

func (g *game) DumpTile(l engine.Letter) {
	g.mu.Lock()
	g.tiles.Dec(l)
	g.mu.Unlock()
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

	// Convert the stgatus to the wire format
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
			tl.Set(i, wordErrs)
		}
		switch status.Code {
		case engine.NotAllLetters:
			resp.SetNotAllLetters(tl)
		case engine.ExtraLetters:
			resp.SetExtraLetters(tl)
		case engine.InvalidWord:
			resp.SetInvalidWord(tl)
		}
	}

	// If the board was valid, increment the players score and let them know waz good
	if status.Code == engine.Success {
		// This can happen in the background, the whole game is asynchronous anyway
		go func() {
			newTiles, err := g.e.addSuccessfulPeel(g.player.name)
			if err != nil {
				log.Println("saving peel: error adding peel and updating tiles: %v", err)
				return
			}

			// TODO(bsprague): Generate potassium.Peel here
			_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
			if err != nil {
				log.Println("saving peel: error making message: %v", err)
				return
			}

			// Create the replay
			p, err := potassium.NewRootPeel(seg)
			if err != nil {
				log.Println("saving peel: error making root peel: %v", err)
				return
			}

			err = p.SetPlayer(g.player.name)
			if err != nil {
				log.Println("saving peel: error setting player name: %v", err)
				return
			}

			err = p.SetBoard(wb)
			if err != nil {
				log.Println("saving peel: error setting board: %v", err)
				return
			}

			tl, err := p.NewNewTiles(int32(len(newTiles)))
			if err != nil {
				log.Println("saving peel: error making new tile list: %v", err)
				return
			}

			i := 0
			for player, letter := range newTiles {
				e := tl.At(i)
				e.SetPlayer(player)
				e.SetLetter(letter.String())
				i++
			}

			err = db.addPeel(p)
			if err != nil {
				log.Println("saving peel: error persisting peel to db: %v", err)
				return
			}
		}()
	}

	return nil
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
	defer g.mu.RUnlock()
	if g.tiles.Freq(letter) <= 0 {
		// They don't actually even have this letter to give away
		resp.SetStatus(potassium.DumpResponse_Status_letterNotInTiles)
		return nil
	}

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
		return nil
	}
	for i, letter := range letters {
		tl.Set(i, letter.String())
	}
	resp.SetStatus(potassium.DumpResponse_Status_success)

	// TODO(bsprague): Write this to the db

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
