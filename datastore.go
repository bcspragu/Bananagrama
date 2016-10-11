package main

import (
	"errors"
	"strconv"
	"sync"
	"time"

	capnp "zombiezen.com/go/capnproto2"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
	"github.com/boltdb/bolt"
)

var (
	errDatastoreNotImplemented = errors.New("bananagrama: datastore operation not implemented")
	errGameNotFound            = errors.New("bananagrama: game not found")
)

type matchID string

type datastore interface {
	// Games
	startGame(playerFreq map[string]engine.FreqList) (matchID, error)
	addPeel(potassium.Peel) error
	addDump(potassium.Dump) error
	finishGame() error
}

// For my own sanity, we store the whole game in-memory until we're ready to
// write it out to disk. Whether or not this is better than
// loading/deserializing/copying/storing every time is up for debate.
type dbImpl struct {
	*bolt.DB
	// fields below are protected by mu
	mu         sync.Mutex
	matchGoing bool
	// The ID to assign to the next match, updated once we persist a game to the
	// DB with finishGame()
	nextID matchID
	game   potassium.Replay
	peels  []potassium.Peel
	dumps  []potassium.Dump
}

var (
	GameBucket = []byte("Games")
)

func initDB(dbName string) (datastore, error) {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	var nextID matchID
	err = db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(GameBucket)
		if err != nil {
			return err
		}

		nID, err := bk.NextSequence()
		if err != nil {
			return err
		}

		nextID = matchID(strconv.FormatUint(nID, 10))

		return nil
	})

	return &dbImpl{
		DB:     db,
		nextID: nextID,
	}, err
}

func (db *dbImpl) startGame(playerFreq map[string]engine.FreqList) (matchID, error) {
	db.mu.Lock()
	if db.matchGoing {
		db.mu.Unlock()
		return "", errors.New("gobots: game in progress, we only do one at a time")
	}
	db.matchGoing = true
	db.mu.Unlock()

	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return "", err
	}

	// Create the replay
	r, err := potassium.NewRootReplay(seg)
	if err != nil {
		return "", err
	}

	// Create the initial tile entries
	it, err := r.NewInitialTiles(int32(len(playerFreq)))
	if err != nil {
		return "", err
	}

	i := 0
	for player, freq := range playerFreq {
		e := it.At(i)
		e.SetPlayer(player)
		il, err := e.NewFrequency(int32(len(freq)))
		if err != nil {
			return "", err
		}

		for j, f := range freq {
			il.Set(j, int32(f))
		}

		i++
	}

	db.game = r
	return db.nextID, nil
}

func (db *dbImpl) addPeel(peel potassium.Peel) error {
	db.peels = append(db.peels, peel)
	return nil
}

func (db *dbImpl) addDump(dump potassium.Dump) error {
	db.dumps = append(db.dumps, dump)
	return nil
}

func (db *dbImpl) finishGame() error {
	pl, err := db.game.NewPeels(int32(len(db.peels)))
	if err != nil {
		return err
	}

	for i, peel := range db.peels {
		err = pl.Set(i, peel)
		if err != nil {
			return err
		}
	}

	dl, err := db.game.NewDumps(int32(len(db.dumps)))
	if err != nil {
		return err
	}

	for i, dump := range db.dumps {
		err = dl.Set(i, dump)
		if err != nil {
			return nil
		}
	}

	err = db.Update(func(tx *bolt.Tx) error {

		return nil
	})

	return err
}
