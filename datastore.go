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
	finishGame(playerScores map[string]int) error

	lookupGame(id matchID) (potassium.Replay, error)
	matchIDs() ([]matchID, error)

	Close() error
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
	msg    *capnp.Message
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
		return "", errors.New("game in progress, we only do one at a time")
	}
	db.matchGoing = true
	db.mu.Unlock()

	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return "", err
	}

	// Create the replay
	r, err := potassium.NewRootReplay(seg)
	if err != nil {
		return "", err
	}

	r.SetStartTime(uint64(time.Now().UnixNano()))

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
	db.msg = msg
	return db.nextID, nil
}

func (db *dbImpl) addPeel(peel potassium.Peel) error {
	peel.SetTimestamp(uint64(time.Now().UnixNano()))
	db.peels = append(db.peels, peel)
	return nil
}

func (db *dbImpl) addDump(dump potassium.Dump) error {
	dump.SetTimestamp(uint64(time.Now().UnixNano()))
	db.dumps = append(db.dumps, dump)
	return nil
}

func (db *dbImpl) finishGame(playerScores map[string]int) error {
	db.mu.Lock()
	if !db.matchGoing {
		db.mu.Unlock()
		return errors.New("game not in progress, certainly can't finish it")
	}
	db.mu.Unlock()

	// Copy over the peels into our persistent data structure
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

	// Copy over the dumps into our persistent data structure
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

	fsl, err := db.game.NewFinalScores(int32(len(playerScores)))
	if err != nil {
		return err
	}

	i := 0
	for player, score := range playerScores {
		rs := fsl.At(i)
		rs.SetPlayer(player)
		rs.SetScore(uint32(score))
		i++
	}

	db.game.SetEndTime(uint64(time.Now().UnixNano()))

	var nextID matchID
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)
		dat, err := db.msg.Marshal()
		if err != nil {
			return err
		}
		err = b.Put([]byte(db.nextID), dat)
		if err != nil {
			return err
		}
		nID, err := b.NextSequence()
		if err != nil {
			return err
		}

		nextID = matchID(strconv.FormatUint(nID, 10))
		return nil
	})

	db.mu.Lock()
	db.matchGoing = false
	db.nextID = nextID
	db.msg = nil
	db.mu.Unlock()
	return err
}

func (db *dbImpl) lookupGame(id matchID) (potassium.Replay, error) {
	var r potassium.Replay
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)
		data := b.Get([]byte(id))
		if len(data) == 0 {
			return errGameNotFound
		}

		msg, err := capnp.Unmarshal(data)
		if err != nil {
			return err
		}
		r, err = potassium.ReadRootReplay(msg)
		return err
	})

	return r, err
}

func (db *dbImpl) matchIDs() ([]matchID, error) {
	var ids []matchID
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(GameBucket)

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			ids = append(ids, matchID(k))
		}

		return nil
	})
	return ids, nil
}
