package banana

import (
	"bytes"
	"fmt"
	"sort"
)

// Board represents an individual players game board.
type Board struct {
	config *boardConfig

	words []Word

	// The below fields are only for caching already-computed state, and are
	// assumed to be in sync with the board.

	// A map from (x,y) to where the letter is on the grid
	letterMap map[Loc]Letter
}

type boardConfig struct {
	minLettersInWord int
}

func (b *boardConfig) Clone() *boardConfig {
	return &boardConfig{
		minLettersInWord: b.minLettersInWord,
	}
}

func NewBoard(c *Config) *Board {
	return &Board{
		config: &boardConfig{
			minLettersInWord: c.MinLettersInWord,
		},
		letterMap: make(map[Loc]Letter),
	}
}

func NewBoardWithWords(c *Config, words []Word) (*Board, error) {
	b := NewBoard(c)
	if err := b.AddWords(words); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Board) Clone() *Board {
	var words []Word
	if b.words != nil {
		words = make([]Word, len(b.words))
		copy(words, b.words)
	}

	letterMap := make(map[Loc]Letter)
	for k, v := range b.letterMap {
		letterMap[k] = v
	}

	return &Board{
		words:     words,
		letterMap: letterMap,
	}
}

func (b *Board) Words() []Word {
	var words []Word
	if b.words != nil {
		words = make([]Word, len(b.words))
		copy(words, b.words)
	}
	return words
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
	return []Loc{
		{l.X + 1, l.Y}, // right
		{l.X - 1, l.Y}, // left
		{l.X, l.Y - 1}, // up
		{l.X, l.Y + 1}, // down
	}
}

func (w Word) CharLocs() []CharLoc {
	dx, dy := 0, 0
	switch w.Orientation {
	case Horizontal:
		dx = 1
	case Vertical:
		dy = 1
	}

	var cls []CharLoc
	for i := 0; i < len(w.Text); i++ {
		if w.Text[i] == ' ' {
			continue
		}
		cls = append(cls, CharLoc{
			Letter: Letter(w.Text[i]),
			Loc:    Loc{X: w.Loc.X + i*dx, Y: w.Loc.Y + i*dy},
		})
	}
	return cls
}

func (b *Board) AddWords(words []Word) error {
	// First, attempt to add all the words to a copy of the board.
	bb := b.Clone()
	for _, w := range words {
		if err := bb.AddWord(w); err != nil {
			return fmt.Errorf("failed to add word to board: %v", err)
		}
	}

	// If we succeeded, we can now add directly to our real board.
	for _, w := range words {
		if err := b.AddWord(w); err != nil {
			// Should never happen, but I'm a realist.
			return fmt.Errorf("failed to add word to board: %v", err)
		}
	}

	return nil
}

// AddWord attempts to add a word to the board, and calculates a few common
// data structures for later Board validation steps. If the word doesn't fit
// with existing words on the board, an error is returned.
func (b *Board) AddWord(w Word) error {
	cls := w.CharLocs()

	// Do one verification loop
	for _, cl := range cls {
		// If we've already placed a letter there, they've sent us a malformed
		// board
		if l, ok := b.letterMap[cl.Loc]; ok && l != cl.Letter {
			return fmt.Errorf("can't place %q in location %+v, %q is already there", cl.Letter, cl.Loc, l)
		}
	}

	// Then do an update loop.
	for _, cl := range cls {
		if _, ok := b.letterMap[cl.Loc]; !ok {
			b.letterMap[cl.Loc] = cl.Letter
		}
	}

	b.words = append(b.words, w)

	return nil
}

func (b *Board) Count() int {
	if b == nil {
		return 0
	}
	return len(b.letterMap)
}

func (b *Board) AsTiles() *Tiles {
	if b == nil {
		return NewTiles()
	}

	t := NewTiles()
	for _, l := range b.letterMap {
		t.Inc(l)
	}

	return t
}

// Validate returns the status of the board.
func (b *Board) Validate(tiles *Tiles, dict Dictionary) *BoardValidation {
	// A board is considered valid if:
	//   - the player used a subset of the letters in their hand
	//   - the words given don't overlap in conflicting ways
	//	 - the words on the grid are real Scrabble words
	//   - the words on the grid are at least the minimum length
	//   - any letter can be reached from any other (aka its all connected)

	unused, unowned := b.leftover(tiles)
	if len(unowned) > 0 {
		return &BoardValidation{ExtraLetters: unowned}
	}

	var invalidWords, shortWords []CharLocs
	for _, cls := range b.findWords() {
		if !dict.HasWord(cls.Word) {
			invalidWords = append(invalidWords, cls)
			continue
		}
		if len(cls.Locs) < b.config.minLettersInWord {
			shortWords = append(shortWords, cls)
		}
	}

	return &BoardValidation{
		InvalidWords:  invalidWords,
		ShortWords:    shortWords,
		DetachedBoard: !b.connected(),
		UnusedLetters: unused,
		ExtraLetters:  unowned,
	}
}

func (b *Board) Diff(tiles *Tiles) *Tiles {
	cp := tiles.Clone()
	for _, letter := range b.letterMap {
		cp.Dec(letter)
	}
	return cp
}

func (b *Board) leftover(tiles *Tiles) (notUsed, notOwned []string) {
	// If diff is greater than zero, it means they didn't use all of that letter in their hand
	// If diff is less than zero, it means they used more than they had
	d := b.Diff(tiles)
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
	d := b.Diff(tiles)
	for _, freq := range d.freq {
		if freq != 0 {
			return false
		}
	}
	return true
}

// findWords builds a list of all the words in the grid
func (b *Board) findWords() []CharLocs {
	usedX := make(map[Loc]Letter)
	usedY := make(map[Loc]Letter)

	// 1. Bucket letters by X, map to lists of (Y, letter)
	xm := make(map[int][]CharLoc)
	for _, word := range b.words {
		for _, cl := range word.CharLocs() {
			if _, ok := usedX[cl.Loc]; ok {
				continue
			}
			xm[cl.Loc.X] = append(xm[cl.Loc.X], cl)
			usedX[cl.Loc] = cl.Letter
		}
	}

	ym := make(map[int][]CharLoc)
	for _, word := range b.words {
		for _, cl := range word.CharLocs() {
			if _, ok := usedY[cl.Loc]; ok {
				continue
			}
			ym[cl.Loc.Y] = append(ym[cl.Loc.Y], cl)
			usedY[cl.Loc] = cl.Letter
		}
	}
	for _, xx := range ym {
		sort.Sort(byX(xx))
	}
	for _, yy := range xm {
		sort.Sort(byY(yy))
	}

	var wds []CharLocs
	// Go down vertically, match up sequential CharLocs into words
	for _, cls := range xm {
		wds = append(wds, findWordsInSequence(cls, func(l Loc) int { return l.Y })...)
	}

	// Go horizontally, match up sequential CharLocs into words
	for _, cls := range ym {
		wds = append(wds, findWordsInSequence(cls, func(l Loc) int { return l.X })...)
	}

	// Doesn't matter if we use usedX or usedY here.
	for loc, l := range usedX {
		if _, ok := usedX[Loc{X: loc.X - 1, Y: loc.Y}]; ok {
			continue
		}
		if _, ok := usedX[Loc{X: loc.X + 1, Y: loc.Y}]; ok {
			continue
		}
		if _, ok := usedX[Loc{X: loc.X, Y: loc.Y - 1}]; ok {
			continue
		}
		if _, ok := usedX[Loc{X: loc.X, Y: loc.Y + 1}]; ok {
			continue
		}
		wds = append(wds, CharLocs{Word: l.String(), Locs: []CharLoc{{Letter: l, Loc: loc}}})
	}

	return wds
}

func (b *Board) connected() bool {
	visitedMap := make(map[Loc]bool)

	// First, populate the visited map.
	for l := range b.letterMap {
		visitedMap[l] = false
	}

	var start Loc
	for loc := range b.letterMap {
		start = loc
		break
	}
	explore(start, visitedMap)

	for _, visited := range visitedMap {
		if !visited {
			return false
		}
	}

	return true
}

func explore(l Loc, visitedMap map[Loc]bool) {
	visitedMap[l] = true
	for _, loc := range l.surrounding() {
		if visited, exists := visitedMap[loc]; exists && !visited {
			explore(loc, visitedMap)
		}
	}
}

func findWordsInSequence(cls []CharLoc, f func(Loc) int) []CharLocs {
	var (
		buf bytes.Buffer
		// Our list of words.
		wds []CharLocs
		// The intermediate list of characters and locations
		tempCls []CharLoc

		started = false
	)

	for i, cur := range cls[:len(cls)-1] {
		next := cls[i+1]

		// Sequential
		if f(next.Loc)-f(cur.Loc) == 1 {
			// Add this letter to the word, and the location to our list.
			buf.WriteRune(rune(cur.Letter))
			tempCls = append(tempCls, cur)

			started = true
			// If we're at the second to last letter, and we know they're sequential,
			// add that one too
			if i == len(cls)-2 {
				buf.WriteRune(rune(next.Letter))
				tempCls = append(tempCls, next)
				wds = append(wds, CharLocs{Word: buf.String(), Locs: tempCls})
			}
		} else if started {
			// If the next two aren't sequential, but we've started a word, that
			// means the current letter is the last one in the word, so we should add
			// it and reset
			buf.WriteRune(rune(cur.Letter))
			tempCls = append(tempCls, cur)
			wds = append(wds, CharLocs{Word: buf.String(), Locs: tempCls})

			buf.Reset()
			tempCls = []CharLoc{}
			started = false
		}
	}
	return wds
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
