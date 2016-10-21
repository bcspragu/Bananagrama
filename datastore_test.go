package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	capnp "zombiezen.com/go/capnproto2"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
)

// newDatastore returns a datastore using a temporary path.
func newDatastore() (datastore, error) {
	// Retrieve a temporary path.
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	path := f.Name()
	f.Close()
	os.Remove(path)
	// Open the database.
	return initDB(path)
}

func freqList(m map[rune]int) engine.FreqList {
	var l engine.FreqList
	for r, c := range m {
		l.Set(engine.Letter(r), c)
	}
	return l
}

func TestDatastore(t *testing.T) {
	db, err := newDatastore()
	if err != nil {
		t.Errorf("error getting db: %v", err)
	}
	defer db.Close()

	id, err := db.startGame(map[string]engine.FreqList{
		"test": freqList(map[rune]int{'a': 1}),
	})

	if err != nil {
		t.Errorf("error starting game: %v", err)
	}

	if id != "1" {
		t.Errorf("matchID: %s, want \"1\"", id)
	}

	// TODO(bsprague): Use makePeel and stuff to test this better
}

func makePeel(player string, board potassium.Board, newTiles []string) (potassium.Peel, error) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return potassium.Peel{}, fmt.Errorf("saving peel: error making message: %v", err)
	}

	p, err := potassium.NewRootPeel(seg)
	if err != nil {
		return potassium.Peel{}, fmt.Errorf("saving peel: error making peel: %v", err)
	}

	bl, err := p.NewValidBoards(1)
	if err != nil {
		return potassium.Peel{}, fmt.Errorf("saving peel: error making board: %v", err)
	}
	b := bl.At(0)
	b.SetPlayer(player)
	b.SetBoard(board)

	nnt, err := p.NewNewTiles(1)
	if err != nil {
		return potassium.Peel{}, fmt.Errorf("saving peel: error making new tiles: %v", err)
	}

	pt := nnt.At(0)
	pt.SetPlayer(player)
	nl, err := pt.NewLetters(int32(len(newTiles)))
	for i, tile := range newTiles {
		nl.Set(i, tile)
	}

	return p, nil
}
