package engine

import (
	"bytes"
	"sort"
)

type Orientation int

const (
	None Orientation = iota
	Horizontal
	Vertical
)

type BoardErrorType int

const (
	Success BoardErrorType = iota
	InvalidWord
	DetachedBoard
)

const (
	CharacterSetSize = 26
	LetterOffset     = 'a'
)

type byX []Word

type Letter rune

func (l Letter) isValid() bool {
	return 'a' <= l && l <= 'z'
}

type FreqList [CharacterSetSize]int

func (f FreqList) GetFreq(l Letter) int {
	if l.isValid() {
		return f[l]
	}
	return 0
}

func (x byX) Len() int { return len(x) }
func (x byX) Less(i, j int) bool {
	if x[i].Loc.X != x[j].Loc.X {
		return x[i].Loc.X < x[j].Loc.X
	}
	return x[i].Loc.Y < x[j].Loc.Y
}
func (x byX) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type byY []Word

func (y byY) Len() int { return len(y) }
func (y byY) Less(i, j int) bool {
	if y[i].Loc.Y != y[j].Loc.Y {
		return y[i].Loc.Y < y[j].Loc.Y
	}
	return y[i].Loc.X < y[j].Loc.X
}
func (y byY) Swap(i, j int) { y[i], y[j] = y[j], y[i] }

type Board struct {
	Words      []Word
	Dictionary Dictionary

	// The below fields are only for caching already-computed state

	// A map from (x,y) to where the letter is on the grid
	letterMap map[Loc]Letter
	// A map from (x,y) to if we've seen a letter or not
	visitedMap map[Loc]bool
}

type Word struct {
	Orientation Orientation
	Text        string
	Loc         Loc
}

type CharLoc struct {
	Letter Letter
	Loc    Loc
}

type Loc struct {
	X, Y int
}

func (l Loc) surrounding() []Loc {
	locs := []Loc{
		{l.X + 1, l.Y},     // right
		{l.X, l.Y + 1},     // down
		{l.X + 1, l.Y + 1}, // down-right
	}

	if l.X > 0 {
		locs = append(locs, Loc{l.X - 1, l.Y} /* left */, Loc{l.X - 1, l.Y + 1} /* down-left */)
	}
	if l.Y > 0 {
		locs = append(locs, Loc{l.X, l.Y - 1} /* up */, Loc{l.X + 1, l.Y - 1} /* up-left*/)
	}

	if l.X > 0 && l.Y > 0 {
		locs = append(locs, Loc{l.X - 1, l.Y - 1}) // up-right
	}
	return locs
}

func (w Word) CharLocs() []CharLoc {
	cls := make([]CharLoc, len(w.Text))
	dx, dy := 0, 0
	switch w.Orientation {
	case Horizontal:
		dx = 1
	case Vertical:
		dy = 1
	}
	for i := 0; i < len(w.Text); i++ {
		cls[i] = CharLoc{
			Letter: Letter(w.Text[i]),
			Loc:    Loc{X: w.Loc.X + i*dx, Y: w.Loc.Y + i*dy},
		}
	}
	return cls
}

// precompute calculates a few common data structures for later Board
// validation steps. If the data structures can't be created, the function
// returns false.
func (b *Board) precompute() bool {
	b.letterMap = make(map[Loc]Letter)
	b.visitedMap = make(map[Loc]bool)

	// Check that the player isn't claiming two different letters occupy the same
	// space
	for _, word := range b.Words {
		for _, cl := range word.CharLocs() {
			// If we've already placed a letter there, they've sent us a malformed
			// board
			if l, ok := b.letterMap[cl.Loc]; ok && l != cl.Letter {
				return false
			} else if !ok {
				b.letterMap[cl.Loc] = cl.Letter
				b.visitedMap[cl.Loc] = false
			}
		}
	}

	return true
}

// TODO(bsprague): Return some error types for invalid boards, oh and lowercase
// words, and check that it only contains alphanumerics, maybe.
func (b *Board) Valid(letters FreqList) bool {
	// A board is considered valid if:
	//   - the player used exactly the letters in their hand
	//   - the words given don't overlap in conflicting ways
	//	 - the words on the grid are real Scrabble words
	//   - any letter can be reached from any other (aka its all connected)

	// If precompute failed, we weren't able to make our intermediate
	// representation, meaning the given words can't form a valid board
	if !b.precompute() {
		return false
	}

	if !b.containsExactly(letters) {
		return false
	}

	// Check for real Scrabble words
	for _, word := range b.findWords() {
		if !b.Dictionary.HasWord(word) {
			return false
		}
	}

	if !b.connected() {
		return false
	}

	return true
}

func (b *Board) containsExactly(letters FreqList) bool {
	cp := letters
	for _, letter := range b.letterMap {
		cp[letter-LetterOffset]--
	}
	for _, freq := range cp {
		if freq != 0 {
			return false
		}
	}
	return true
}

// findWords builds a list of all the words in the grid
func (b *Board) findWords() []string {
	strs := []string{}
	// The algorithm:
	// 1. Bucket letters by X, map to lists of (Y, letter)
	xm := make(map[int][]CharLoc)
	sort.Sort(byY(b.Words))
	for _, word := range b.Words {
		for _, cl := range word.CharLocs() {
			xm[cl.Loc.X] = append(xm[cl.Loc.X], cl)
		}
	}

	ym := make(map[int][]CharLoc)
	sort.Sort(byX(b.Words))
	for _, word := range b.Words {
		for _, cl := range word.CharLocs() {
			ym[cl.Loc.Y] = append(ym[cl.Loc.Y], cl)
		}
	}

	// Go down vertically, match up sequential CharLocs into words
	for _, cls := range xm {
		strs = append(strs, findWordsInSequence(cls, func(l Loc) int { return l.Y })...)
	}

	// Go horizontally, match up sequential CharLocs into words
	for _, cls := range ym {
		strs = append(strs, findWordsInSequence(cls, func(l Loc) int { return l.X })...)
	}

	return strs
}

func (b *Board) connected() bool {
	var start Loc
	for loc := range b.letterMap {
		start = loc
		break
	}
	b.explore(start)

	for _, visited := range b.visitedMap {
		if !visited {
			return false
		}
	}

	return true
}

func (b *Board) explore(l Loc) {
	b.visitedMap[l] = true
	for _, loc := range l.surrounding() {
		if visited, exists := b.visitedMap[loc]; exists && !visited {
			b.explore(loc)
		}
	}
}

func findWordsInSequence(cls []CharLoc, f func(Loc) int) []string {
	var buf bytes.Buffer
	strs := []string{}
	started := false
	for i, cur := range cls[:len(cls)-1] {
		next := cls[i+1]
		// Sequential
		if f(next.Loc)-f(cur.Loc) == 1 {
			// Add this letter to the word
			buf.WriteRune(rune(cur.Letter))
			started = true
			// If we're at the second to last letter, and we know they're sequential,
			// add that one too
			if i == len(cls)-2 {
				buf.WriteRune(rune(next.Letter))
				strs = append(strs, buf.String())
			}
		} else if started {
			// If the next two aren't sequential, but we've started a word, that
			// means the current letter is the last one in the word, so we should add
			// it and reset
			buf.WriteRune(rune(cur.Letter))
			strs = append(strs, buf.String())
			buf.Reset()
			started = false
		}
	}
	return strs
}
