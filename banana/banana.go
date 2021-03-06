// Package banana contains the domain types for playing a game of Bananagrams.
package banana

import "time"

type GameID string

type DB interface {
	// Creates a new game with the given name and creator.
	NewGame(name string, creator PlayerID, config *Config) (GameID, error)
	// Get all of the games.
	Games() ([]*Game, error)
	// Loads a game with the given ID.
	Game(id GameID) (*Game, error)
	// Loads the bunch for the game with the given ID.
	Bunch(id GameID) (*Bunch, error)
	// Registers a player in our system.
	RegisterPlayer(name string) (PlayerID, error)
	// Adds a player to a not-yet-started game.
	AddPlayerToGame(gID GameID, pID PlayerID) error
	// Get all the players for a game.
	Players(id GameID) ([]*Player, error)
	// Loads a player with the given ID.
	Player(id PlayerID) (*Player, error)
	// Loads the board for the given game and player IDs.
	Board(gID GameID, pID PlayerID) (*Board, error)
	// Loads tiles for the given game and player IDs.
	Tiles(gID GameID, pID PlayerID) (*Tiles, error)
	// Updates a player's board.
	UpdateBoard(gID GameID, pID PlayerID, board *Board) error
	// Updates a player's tiles.
	UpdateTiles(gID GameID, pID PlayerID, tiles *Tiles) error
	// Updates the bunch for the game.
	UpdateBunch(id GameID, bunch *Bunch) error
	// Starts a game, and sets everyone's initial tile sets.
	StartGame(id GameID, players map[PlayerID]*Tiles, bunch *Bunch) error
	// Ends a given game, stops players from adding more to their boards.
	EndGame(id GameID) error
}

type GameStatus int

const (
	UnknownStatus GameStatus = iota
	WaitingForPlayers
	InProgress
	Finished
)

func (g GameStatus) String() string {
	switch g {
	case WaitingForPlayers:
		return "Waiting For Players"
	case InProgress:
		return "In Progress"
	case Finished:
		return "Finished"
	default:
		return "Unknown"
	}
}

type Game struct {
	ID        GameID
	Creator   PlayerID
	Name      string
	Status    GameStatus
	CreatedAt time.Time
	Config    *Config
}

type Config struct {
	// The minimum number of letters that a word needs to have to be considered
	// valid.
	MinLettersInWord int
}

func (c *Config) Clone() *Config {
	return &Config{
		MinLettersInWord: c.MinLettersInWord,
	}
}

func (g *Game) Clone() *Game {
	return &Game{
		ID:        g.ID,
		Creator:   g.Creator,
		Name:      g.Name,
		Status:    g.Status,
		CreatedAt: g.CreatedAt,
		Config:    g.Config.Clone(),
	}
}

type PlayerID string

type Player struct {
	ID      PlayerID
	Name    string
	AddedAt time.Time
}

func (p *Player) Clone() *Player {
	return &Player{
		ID:      p.ID,
		Name:    p.Name,
		AddedAt: p.AddedAt,
	}
}

// Orientation describes how a board is placed on the board.
type Orientation int

const (
	// NoOrientation is a catch-all for unknown orientations.
	NoOrientation Orientation = iota
	// Horizontal means the word is placed on the board from left to right.
	Horizontal
	// Vertical means the word is placed on the board from top to bottom.
	Vertical
)

type BoardValidation struct {
	InvalidWords  []CharLocs
	ShortWords    []CharLocs
	InvalidBoard  bool
	DetachedBoard bool
	UnusedLetters []string
	ExtraLetters  []string
}

// CharLocs is a list of letters and the word they
// make up, this is currently only used for
// returning bad words.
type CharLocs struct {
	// The word that is made of the charlocs.
	Word string
	Locs []CharLoc
}

// BoardStatus describes if a board is valid, or how it is invalid.
type BoardStatus struct {
	Code   BoardStatusCode
	Errors []string
}

// BoardStatusCode describes the current validity of a board.
type BoardStatusCode int

const (
	Success BoardStatusCode = iota
	InvalidWord
	DetachedBoard
	NotAllLetters
	ExtraLetters
	InvalidBoard
)

// byX is a way of sorting a slice of charlocs that sorts them first according to
// their X position, then by their Y position.
type byX []CharLoc

func (x byX) Len() int { return len(x) }
func (x byX) Less(i, j int) bool {
	if x[i].Loc.X != x[j].Loc.X {
		return x[i].Loc.X < x[j].Loc.X
	}
	return x[i].Loc.Y < x[j].Loc.Y
}
func (x byX) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// byY is a way of sorting a slice of charlocs that sorts them first according to
// their Y position, then by their X position.
type byY []CharLoc

func (y byY) Len() int { return len(y) }
func (y byY) Less(i, j int) bool {
	if y[i].Loc.Y != y[j].Loc.Y {
		return y[i].Loc.Y < y[j].Loc.Y
	}
	return y[i].Loc.X < y[j].Loc.X
}
func (y byY) Swap(i, j int) { y[i], y[j] = y[j], y[i] }
