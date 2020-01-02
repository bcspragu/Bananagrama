package memdb

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/bcspragu/Bananagrama/banana"
)

var (
	ErrGameNotFound   = errors.New("game not found")
	ErrGameNotStarted = errors.New("game not started yet")
	ErrPlayerNotFound = errors.New("player not found")
)

type DB struct {
	r   *rand.Rand
	now func() time.Time

	sync.RWMutex
	// Game-keyed maps
	games         map[banana.GameID]*banana.Game
	gameToPlayers map[banana.GameID][]banana.PlayerID
	bunches       map[banana.GameID]*banana.Bunch

	// Player-keyed maps
	players map[banana.PlayerID]*banana.Player
	boards  map[banana.PlayerID]*banana.Board
	tiles   map[banana.PlayerID]*banana.Tiles
}

func New(r *rand.Rand) *DB {
	return &DB{
		games:         make(map[banana.GameID]*banana.Game),
		gameToPlayers: make(map[banana.GameID][]banana.PlayerID),
		bunches:       make(map[banana.GameID]*banana.Bunch),

		players: make(map[banana.PlayerID]*banana.Player),
		boards:  make(map[banana.PlayerID]*banana.Board),
		tiles:   make(map[banana.PlayerID]*banana.Tiles),

		r:   r,
		now: time.Now,
	}
}

// Creates a new game with the given name.
func (d *DB) NewGame(name string, creator banana.PlayerID) (banana.GameID, error) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 10; i++ {
		gID := banana.GameID(d.randomID())
		if _, ok := d.games[gID]; ok {
			// If this ID is taken, that's really unlikely and strange, but just try
			// again.
			continue
		}
		d.games[gID] = &banana.Game{
			ID:        gID,
			Creator:   creator,
			Name:      name,
			Status:    banana.WaitingForPlayers,
			CreatedAt: time.Now(),
		}
		d.gameToPlayers[gID] = []banana.PlayerID{}
		return gID, nil
	}
	return banana.GameID(""), errors.New("failed to find unique game ID after 10 tries, something is terribly wrong")
}

func (d *DB) Games() ([]*banana.Game, error) {
	d.RLock()
	defer d.RUnlock()

	var gs []*banana.Game
	for _, g := range d.games {
		gs = append(gs, g.Clone())
	}

	sort.Slice(gs, func(i, j int) bool {
		return gs[i].CreatedAt.Before(gs[j].CreatedAt)
	})

	return gs, nil
}

// Loads a game with the given ID.
func (d *DB) Game(id banana.GameID) (*banana.Game, error) {
	d.RLock()
	defer d.RUnlock()

	g, ok := d.games[id]
	if !ok {
		return nil, ErrGameNotFound
	}
	return g.Clone(), nil
}

// Loads the bunch for the game with the given ID.
func (d *DB) Bunch(id banana.GameID) (*banana.Bunch, error) {
	d.RLock()
	defer d.RUnlock()

	if _, ok := d.games[id]; !ok {
		return nil, ErrGameNotFound
	}

	b, ok := d.bunches[id]
	if !ok {
		return nil, ErrGameNotStarted
	}
	return b.Clone(), nil
}

// Adds a player to a game.
func (d *DB) AddPlayer(gID banana.GameID, name string) (banana.PlayerID, error) {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.games[gID]; !ok {
		return banana.PlayerID(""), ErrGameNotFound
	}

	for i := 0; i < 10; i++ {
		pID := banana.PlayerID(d.randomID())
		if _, ok := d.players[pID]; ok {
			// If this ID is taken, that's really unlikely and strange, but just try
			// again.
			continue
		}

		// If we're here, we've found a unique ID that we can use and return.
		p := &banana.Player{
			ID:      pID,
			Name:    name,
			AddedAt: d.now(),
		}
		d.gameToPlayers[gID] = append(d.gameToPlayers[gID], pID)
		d.players[pID] = p
		d.boards[pID] = &banana.Board{}
		d.tiles[pID] = banana.NewTiles()

		return pID, nil
	}

	return banana.PlayerID(""), errors.New("failed to find unique player ID after 10 tries, something is terribly wrong")
}

// Get all the players for a given game.
func (d *DB) Players(id banana.GameID) ([]*banana.Player, error) {
	d.RLock()
	defer d.RUnlock()

	pIDs, ok := d.gameToPlayers[id]
	if !ok {
		return nil, ErrGameNotFound
	}

	var ps []*banana.Player
	for _, pID := range pIDs {
		ps = append(ps, d.players[pID].Clone())
	}

	return ps, nil
}

// Loads a player with the given ID.
func (d *DB) Player(pID banana.PlayerID) (*banana.Player, error) {
	d.RLock()
	defer d.RUnlock()

	p, ok := d.players[pID]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return p.Clone(), nil
}

func (d *DB) Board(pID banana.PlayerID) (*banana.Board, error) {
	d.RLock()
	defer d.RUnlock()

	b, ok := d.boards[pID]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return b.Clone(), nil
}

func (d *DB) Tiles(pID banana.PlayerID) (*banana.Tiles, error) {
	d.RLock()
	defer d.RUnlock()

	t, ok := d.tiles[pID]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return t.Clone(), nil
}

// Updates a player's board.
func (d *DB) UpdateBoard(pID banana.PlayerID, board *banana.Board) error {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.boards[pID]; !ok {
		return ErrPlayerNotFound
	}

	d.boards[pID] = board
	return nil
}

// Updates a player's tiles.
func (d *DB) UpdateTiles(pID banana.PlayerID, tiles *banana.Tiles) error {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.tiles[pID]; !ok {
		return ErrPlayerNotFound
	}

	d.tiles[pID] = tiles
	return nil
}

// Updates the bunch for the game.
func (d *DB) UpdateBunch(gID banana.GameID, bunch *banana.Bunch) error {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.bunches[gID]; !ok {
		return ErrGameNotFound
	}

	d.bunches[gID] = bunch
	return nil
}

// Starts a game, and sets everyone's initial tile sets.
func (d *DB) StartGame(gID banana.GameID, players map[banana.PlayerID]*banana.Tiles, bunch *banana.Bunch) error {
	d.Lock()
	defer d.Unlock()

	g, ok := d.games[gID]
	if !ok {
		return ErrGameNotFound
	}
	g.Status = banana.InProgress
	d.bunches[gID] = bunch.Clone()

	if n := len(d.gameToPlayers[gID]); n != len(players) {
		return fmt.Errorf("%d players in game, %d players in tile set map", n, len(players))
	}

	for pID, tiles := range players {
		if _, ok := d.tiles[pID]; !ok {
			return ErrPlayerNotFound
		}
		d.tiles[pID] = tiles.Clone()
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
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[d.r.Intn(len(letters))]
	}
	return string(b)
}
