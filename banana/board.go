package banana

import (
	"bytes"
	"sort"
)

// Board represents an individual players game board.
type Board struct {
	Words      []Word
	Dictionary Dictionary

	// The below fields are only for caching already-computed state

	// A map from (x,y) to where the letter is on the grid
	letterMap map[Loc]Letter
	// A map from (x,y) to if we've seen a letter or not
	visitedMap map[Loc]bool
}

func (b *Board) clone() *Board {
	var (
		words      []Word
		letterMap  map[Loc]Letter
		visitedMap map[Loc]bool
	)

	if b.Words != nil {
		words = make([]Word, len(b.Words))
		copy(words, b.Words)
	}

	if b.letterMap != nil {
		letterMap = make(map[Loc]Letter)
		for k, v := range b.letterMap {
			letterMap[k] = v
		}
	}

	if b.visitedMap != nil {
		visitedMap = make(map[Loc]bool)
		for k, v := range b.visitedMap {
			visitedMap[k] = v
		}
	}

	return &Board{
		Words:      words,
		Dictionary: b.Dictionary,
		letterMap:  letterMap,
		visitedMap: visitedMap,
	}
}

// Word represents the placement of a single word on a Bananagrams board.
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

// TODO(bsprague): Lowercase words, and check that it only contains
// alphanumerics, maybe.
func (b *Board) ValidateBoard(tiles *Tiles) BoardStatus {
	// A board is considered valid if:
	//   - the player used exactly the letters in their hand
	//   - the words given don't overlap in conflicting ways
	//	 - the words on the grid are real Scrabble words
	//   - any letter can be reached from any other (aka its all connected)

	// If precompute failed, we weren't able to make our intermediate
	// representation, meaning the given words can't form a valid board
	if !b.precompute() {
		return BoardStatus{Code: InvalidBoard}
	}

	unused, unowned := b.leftover(tiles)
	if len(unused) > 0 {
		return BoardStatus{Code: NotAllLetters, Errors: unused}
	}

	if len(unowned) > 0 {
		return BoardStatus{Code: ExtraLetters, Errors: unowned}
	}

	// Check for real Scrabble words
	for _, word := range b.findWords() {
		if !b.Dictionary.HasWord(word) {
			return BoardStatus{Code: InvalidWord, Errors: []string{word}}
		}
	}

	if !b.connected() {
		return BoardStatus{Code: DetachedBoard}
	}

	return BoardStatus{Code: Success}
}

func (b *Board) diff(tiles *Tiles) *Tiles {
	cp := tiles.clone()
	for _, letter := range b.letterMap {
		cp.Dec(letter)
	}
	return cp
}

func (b *Board) leftover(tiles *Tiles) (notUsed, notOwned []string) {
	// If diff is greater than zero, it means they didn't use all of that letter in their hand
	// If diff is less than zero, it means they used more than they had
	d := b.diff(tiles)
	for l, freq := range d.freq {
		if freq > 0 {
			notUsed = append(notUsed, l.String())
			continue
		}
		if freq < 0 {
			notOwned = append(notOwned, l.String())
			continue
		}
	}
	sort.Slice(notUsed, func(i, j int) bool { return notUsed[i] < notUsed[j] })
	sort.Slice(notOwned, func(i, j int) bool { return notOwned[i] < notOwned[j] })
	return notUsed, notOwned
}

func (b *Board) containsExactly(tiles *Tiles) bool {
	// TODO: This could probably be written more directly by just comparing
	// b.letterMap and tiles.
	d := b.diff(tiles)
	for _, freq := range d.freq {
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

func StartingTileCount(pc, scale int) int {
	var tc int
	switch pc {
	case 1, 2, 3, 4:
		tc = 21
	case 5, 6:
		tc = 15
	case 7, 8:
		tc = 11
	default:
		// Default to 11, if we for some reason have more than 8 people...or zero.
		tc = 11
	}
	return tc * scale
}
