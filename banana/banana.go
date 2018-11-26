// Package banana contains the domain types for playing a game of Bananagrams.
package banana

type GameID string

type DB interface {
	// Creates a new game with the given name.
	NewGame(name string) (GameID, error)
	// Loads a game with the given ID.
	Game(id GameID) (*Game, error)
	// Adds a player to a not-yet-started game.
	AddPlayer(id GameID, name string) (PlayerID, error)
	// Updates a player's board.
	UpdatePlayer(id PlayerID, board *Board, tiles *Tiles) error
	// Starts a game, and sets everyone's initial tile sets.
	StartGame(id GameID, players map[PlayerID]*Tiles) error
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

type Game struct {
	Name    string
	Players []*Player
	Status  GameStatus
}

type PlayerID string

type Player struct {
	ID    PlayerID
	Name  string
	Board *Board
	Tiles *Tiles
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

// byX is a way of sorting a slice of words that sorts them first according to
// their X position, then by their Y position.
type byX []Word

func (x byX) Len() int { return len(x) }
func (x byX) Less(i, j int) bool {
	if x[i].Loc.X != x[j].Loc.X {
		return x[i].Loc.X < x[j].Loc.X
	}
	return x[i].Loc.Y < x[j].Loc.Y
}
func (x byX) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// byY is a way of sorting a slice of words that sorts them first according to
// their Y position, then by their X position.
type byY []Word

func (y byY) Len() int { return len(y) }
func (y byY) Less(i, j int) bool {
	if y[i].Loc.Y != y[j].Loc.Y {
		return y[i].Loc.Y < y[j].Loc.Y
	}
	return y[i].Loc.X < y[j].Loc.X
}
func (y byY) Swap(i, j int) { y[i], y[j] = y[j], y[i] }
