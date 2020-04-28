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
	ErrNameTaken      = errors.New("player name already taken")
)

type DB struct {
	r   *rand.Rand
	now func() time.Time

	sync.RWMutex
	// Game-keyed maps
	games         map[banana.GameID]*banana.Game
	gameToPlayers map[banana.GameID][]banana.PlayerID
	bunches       map[banana.GameID]*banana.Bunch
	boards        map[banana.GameID]map[banana.PlayerID]*banana.Board
	tiles         map[banana.GameID]map[banana.PlayerID]*banana.Tiles

	// Player-keyed maps
	players     map[banana.PlayerID]*banana.Player
	playerNames map[string]banana.PlayerID
}

func New(r *rand.Rand) *DB {
	return &DB{
		games:         make(map[banana.GameID]*banana.Game),
		gameToPlayers: make(map[banana.GameID][]banana.PlayerID),
		bunches:       make(map[banana.GameID]*banana.Bunch),
		boards:        make(map[banana.GameID]map[banana.PlayerID]*banana.Board),
		tiles:         make(map[banana.GameID]map[banana.PlayerID]*banana.Tiles),

		players:     make(map[banana.PlayerID]*banana.Player),
		playerNames: make(map[string]banana.PlayerID),

		r:   r,
		now: time.Now,
	}
}

// Creates a new game with the given name.
func (d *DB) NewGame(name string, creator banana.PlayerID, config *banana.Config) (banana.GameID, error) {
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
			Config:    config.Clone(),
		}
		d.gameToPlayers[gID] = []banana.PlayerID{}
		d.boards[gID] = make(map[banana.PlayerID]*banana.Board)
		d.tiles[gID] = make(map[banana.PlayerID]*banana.Tiles)
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

func (d *DB) RegisterPlayer(name string) (banana.PlayerID, error) {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.playerNames[name]; ok {
		return "", ErrNameTaken
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
		d.players[pID] = p
		d.playerNames[name] = pID

		return pID, nil
	}

	return banana.PlayerID(""), errors.New("failed to find unique player ID after 10 tries, something is terribly wrong")
}

// Adds a player to a game.
func (d *DB) AddPlayerToGame(gID banana.GameID, pID banana.PlayerID) error {
	d.Lock()
	defer d.Unlock()

	g, ok := d.games[gID]
	if !ok {
		return ErrGameNotFound
	}

	if _, ok := d.players[pID]; !ok {
		return ErrPlayerNotFound
	}

	d.gameToPlayers[gID] = append(d.gameToPlayers[gID], pID)
	d.boards[gID][pID] = banana.NewBoard(g.Config)
	d.tiles[gID][pID] = banana.NewTiles()

	return nil
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

func (d *DB) Board(gID banana.GameID, pID banana.PlayerID) (*banana.Board, error) {
	d.RLock()
	defer d.RUnlock()

	gm, ok := d.boards[gID]
	if !ok {
		return nil, ErrGameNotFound
	}

	b, ok := gm[pID]
	if !ok {
		return nil, ErrPlayerNotFound
	}

	return b.Clone(), nil
}

func (d *DB) Tiles(gID banana.GameID, pID banana.PlayerID) (*banana.Tiles, error) {
	d.RLock()
	defer d.RUnlock()

	gm, ok := d.tiles[gID]
	if !ok {
		return nil, ErrGameNotFound
	}

	t, ok := gm[pID]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return t.Clone(), nil
}

// Updates a player's board.
func (d *DB) UpdateBoard(gID banana.GameID, pID banana.PlayerID, board *banana.Board) error {
	d.Lock()
	defer d.Unlock()

	gm, ok := d.boards[gID]
	if !ok {
		return ErrGameNotFound
	}

	if _, ok := gm[pID]; !ok {
		return ErrPlayerNotFound
	}

	gm[pID] = board.Clone()

	return nil
}

// Updates a player's tiles.
func (d *DB) UpdateTiles(gID banana.GameID, pID banana.PlayerID, tiles *banana.Tiles) error {
	d.Lock()
	defer d.Unlock()

	gm, ok := d.tiles[gID]
	if !ok {
		return ErrGameNotFound
	}

	if _, ok := gm[pID]; !ok {
		return ErrPlayerNotFound
	}

	gm[pID] = tiles.Clone()
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
		if _, ok := d.tiles[gID][pID]; !ok {
			return ErrPlayerNotFound
		}
		d.tiles[gID][pID] = tiles.Clone()
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
