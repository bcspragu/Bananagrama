package main

import (
	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
)

type game struct {
	bunch *engine.Bunch
	tiles engine.FreqList
}

func (g *game) Peel(call potassium.Game_peel) error {
	return nil
}

func (g *game) Dump(call potassium.Game_peel) error {
	return nil
}

func boardFromWire(b potassium.Board) (*engine.Board, error) {
	// TODO(bsprague): Implement plz
	//wordList, err := b.Words()
	//if err != nil {
	//return nil, err
	//}

	return nil, nil
}

func wordFromWire(w potassium.Word) (engine.Word, error) {
	return engine.Word{}, nil
}
