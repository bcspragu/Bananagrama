package memdb

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"github.com/bcspragu/Bananagrama/banana"
)

var (
	ErrGameNotFound   = errors.New("game not found")
	ErrPlayerNotFound = errors.New("player not found")
)

type DB struct {
	sync.RWMutex
	dict    banana.Dictionary
	games   map[banana.GameID]*banana.Game
	players map[banana.PlayerID]*banana.Player
	r       *rand.Rand
}

func New(r *rand.Rand, dict banana.Dictionary) *DB {
	return &DB{
		dict:    dict,
		games:   make(map[banana.GameID]*banana.Game),
		players: make(map[banana.PlayerID]*banana.Player),
		r:       r,
	}
}

// Creates a new game with the given name.
func (d *DB) NewGame(name string) (banana.GameID, error) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 10; i++ {
		id := banana.GameID(d.randomID())
		if _, ok := d.games[id]; ok {
			// If this ID is taken, that's really unlikely and strange, but just try
			// again.
			continue
		}
		d.games[id] = &banana.Game{
			Name:    name,
			Players: []*banana.Player{},
			Status:  banana.WaitingForPlayers,
		}
		return id, nil
	}
	return banana.GameID(""), errors.New("failed to find unique game ID after 10 tries, something is terribly wrong")
}

// Loads a game with the given ID.
func (d *DB) Game(id banana.GameID) (*banana.Game, error) {
	d.RLock()
	defer d.RUnlock()

	g, ok := d.games[id]
	if !ok {
		return nil, ErrGameNotFound
	}
	return g, nil
}

// Adds a player to a not-yet-started game.
func (d *DB) AddPlayer(id banana.GameID, name string) (banana.PlayerID, error) {
	d.Lock()
	defer d.Unlock()

	g, ok := d.games[id]
	if !ok {
		return banana.PlayerID(""), ErrGameNotFound
	}

	pid, err := d.randomPlayerID(id)
	if err != nil {
		return banana.PlayerID(""), err
	}

	p := &banana.Player{
		ID:   pid,
		Name: name,
		Board: &banana.Board{
			Dictionary: d.dict,
		},
	}

	g.Players = append(g.Players, p)
	d.players[pid] = p

	return pid, nil
}

// randomPlayerID returns a random player ID that includes the given GameID.
// This function expects that the caller already holds a lock on the internal
// state of the DB.
func (d *DB) randomPlayerID(id banana.GameID) (banana.PlayerID, error) {
	for i := 0; i < 10; i++ {
		pid := banana.PlayerID(string(id) + ":" + d.randomID())
		if _, ok := d.players[pid]; ok {
			// If this ID is taken, that's really unlikely and strange, but just try
			// again.
			continue
		}
		// If we're here, we've found a unique ID that we can return.
		return pid, nil
	}
	return banana.PlayerID(""), errors.New("failed to find unique player ID after 10 tries, something is terribly wrong")
}

// Updates a player's board.
func (d *DB) UpdatePlayer(id banana.PlayerID, board *banana.Board, tiles *banana.Tiles) error {
	d.Lock()
	defer d.Unlock()

	p, ok := d.players[id]
	if !ok {
		return ErrPlayerNotFound
	}

	p.Board = board
	p.Tiles = tiles

	return nil
}

// Starts a game, and sets everyone's initial tile sets.
func (d *DB) StartGame(id banana.GameID, players map[banana.PlayerID]*banana.Tiles) error {
	d.Lock()
	defer d.Unlock()

	g, ok := d.games[id]
	if !ok {
		return ErrGameNotFound
	}
	g.Status = banana.InProgress

	if len(g.Players) != len(players) {
		return fmt.Errorf("%d players in game, %d players in tile set map", len(g.Players), len(players))
	}

	for pid, tiles := range players {
		p, ok := d.players[pid]
		if !ok {
			return ErrPlayerNotFound
		}
		p.Tiles = tiles
	}

	return nil
}

// Ends a given game, stops players from adding more to their boards.
func (d *DB) EndGame(id banana.GameID) error {
	d.Lock()
	defer d.Unlock()

	g, ok := d.games[id]
	if !ok {
		return ErrGameNotFound
	}

	g.Status = banana.Finished
	return nil
}

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (d *DB) randomID() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = letters[d.r.Intn(len(letters))]
	}
	return string(b)
}
