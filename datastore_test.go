package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bcspragu/Bananagrama/engine"
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

	// TODO(bsprague): Test other things, like adding peels and dups and
	// finishing the game and looking one up
}
