package main

import (
	"errors"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

var (
	errDatastoreNotImplemented = errors.New("bananagrama: datastore operation not implemented")
	errGameNotFound            = errors.New("bananagrama: game not found")
)

type datastore interface {
	// TODO(bsprague): Decide on database API
}

type dbImpl struct {
	*bolt.DB
	// fields below are protected by mu
	mu sync.Mutex
}

var (
	GameBucket = []byte("Games")
)

func initDB(dbName string) (datastore, error) {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		for _, b := range [][]byte{GameBucket} {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}

		return nil
	})

	return &dbImpl{
		DB: db,
	}, err
}
