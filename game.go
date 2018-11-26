package main

import (
	"fmt"
	"sync"

	"github.com/bcspragu/Bananagrama/banana"
	capnp "zombiezen.com/go/capnproto2"
)

var (
	wireOrientationMap = map[pb.Word_Orientation]banana.Orientation{
		pb.Word_UNKNOWN:    banana.None,
		pb.Word_HORIZONTAL: banana.Horizontal,
		pb.Word_VERTICAL:   banana.Vertical,
	}

	engineStatusMap = map[banana.BoardStatusCode]pb.PeelResponse_Status{
		banana.Success:       pb.PeelResponse_SUCCESS,
		banana.InvalidWord:   pb.PeelResponse_INVALID_WORD,
		banana.DetachedBoard: pb.PeelResponse_DETACHED_BOARD,
		banana.NotAllLetters: pb.PeelResponse_NOT_ALL_LETTERS,
		banana.ExtraLetters:  pb.PeelResponse_EXTRA_LETTERS,
	}
)

type player struct {
	name string
}

type game struct {
	e *aiEndpoint

	mu     sync.RWMutex
	score  int
	tiles  banana.Tiles
	player *player
}

func (g *game) Peel(call potassium.Game_peel) error {
}

type peelInfo struct {
	player string
	board  potassium.Board
}

// Dump exchanges a player's letter for three from the bunch
func (g *game) Dump(call potassium.Game_dump) error {
	g.e.mu.Lock()
	defer g.e.mu.Unlock()

	req := call.Params
	resp := call.Results

	if !g.e.gameStarted {
		resp.SetStatus(potassium.DumpResponse_Status_gameNotStarted)
		return nil
	}

	// The letter the player wants to dump
	l, err := req.Letter()
	if err != nil {
		return err
	}

	if len(l) != 1 {
		resp.SetStatus(potassium.DumpResponse_Status_malformedRequest)
		return nil
	}

	letter := banana.Letter(l[0])

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

func boardFromWire(b potassium.Board) (*banana.Board, error) {
	wireWords, err := b.Words()
	if err != nil {
		return nil, err
	}

	words := make([]banana.Word, wireWords.Len())
	for i := 0; i < wireWords.Len(); i++ {
		w, err := wordFromWire(wireWords.At(i))
		if err != nil {
			return nil, err
		}
		words[i] = w
	}

	return &banana.Board{
		Words:      words,
		Dictionary: dict,
	}, nil
}

func wordFromWire(w potassium.Word) (banana.Word, error) {
	text, err := w.Text()
	if err != nil {
		return banana.Word{}, nil
	}

	return banana.Word{
		Orientation: wireOrientationMap[w.Orientation()],
		Text:        text,
		Loc:         banana.Loc{X: int(w.X()), Y: int(w.Y())},
	}, nil
}
