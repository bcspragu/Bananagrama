package engine

import "sort"

type Orientation int

const (
	None Orientation = iota
	Horizontal
	Vertical
)

type byX []Word

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
	Words []Word
}

type Word struct {
	Orientation Orientation
	Text        string
	Loc         Loc
}

type CharLoc struct {
	Letter string
	Loc    Loc
}

type Loc struct {
	X, Y int
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
			Letter: w.Text[i : i+1],
			Loc:    Loc{X: w.Loc.X + i*dx, Y: w.Loc.Y + i*dy},
		}
	}
	return cls
}

// TODO(bsprague): Return some error types for invalid boards
func (b *Board) Valid(letters []string) bool {
	letterMap := make(map[Loc]string)

	for _, word := range b.Words {
		for _, cl := range word.CharLocs() {
			// If we've already placed a letter there, they've sent us a malformed
			// board
			if l, ok := letterMap[cl.Loc]; ok && l != cl.Letter {
				return false
			} else if !ok {
				letterMap[cl.Loc] = cl.Letter
			}
		}
	}
	// TODO(bsprague): Check that they're all real words, and check for words created by juxtaposition
	return true
}

type other struct {
	Coor   int
	Letter string
}

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

func findWordsInSequence(cls []CharLoc, f func(Loc) int) []string {
	strs := []string{}
	started := false
	w := ""
	for i, cur := range cls[:len(cls)-1] {
		next := cls[i+1]
		// Sequential
		if f(next.Loc)-f(cur.Loc) == 1 {
			// Add this letter to the word
			w += cur.Letter
			started = true
			if i == len(cls)-2 {
				w += next.Letter
			}
		} else {
			// If the next two aren't sequential, but we've started a word, that
			// means the current letter is the last one in the word, so we should add
			// it and reset
			if started {
				w += cur.Letter
				strs = append(strs, w)
				w = ""
				started = false
			}
		}
	}
	if w != "" {
		strs = append(strs, w)
	}
	return strs
}
