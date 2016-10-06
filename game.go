package main

import (
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

type game struct {
	e *aiEndpoint
	// TODO(bsprague): Synchronize on tiles
	tiles  engine.FreqList
	player potassium.Player
	name   string
}

func (g *game) Peel(call potassium.Game_peel) error {
	req := call.Params
	resp := call.Results

	wb, err := req.Board()
	if err != nil {
		return err
	}

	b, err := boardFromWire(wb)
	if err != nil {
		return err
	}

	status := b.Status(g.tiles)
	if status.Code == engine.Success {
		g.e.sendNewTiles(g.name)
	}
	resp.SetStatus(engineStatusMap[status.Code])
	// TODO(bsprague): Check more errors
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

	return nil
}

func (g *game) Dump(call potassium.Game_dump) error {
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
