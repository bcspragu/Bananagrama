// Package banana contains the domain types for playing a game of Bananagrams.
package banana

type DB interface {
	// Games
	StartGame(players map[string]*Tiles) (matchID, error)
	AddPeel(potassium.Peel) error
	AddDump(potassium.Dump) error
	finishGame(playerScores map[string]int) error

	lookupGame(id matchID) (potassium.Replay, error)
	matchIDs() ([]matchID, error)

	Close() error
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
