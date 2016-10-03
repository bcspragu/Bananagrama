package engine

import "sort"

type Orientation int

const (
	None Orientation = iota
	Horizontal
	Vertical
)

type byX []Word

func (x byX) Len() int           { return len(x) }
func (x byX) Less(i, j int) bool { return x[i].Loc.X < x[j].Loc.Y }
func (x byX) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byY []Word

func (y byY) Len() int           { return len(y) }
func (y byY) Less(i, j int) bool { return y[i].Loc.Y < y[j].Loc.Y }
func (y byY) Swap(i, j int)      { y[i], y[j] = y[j], y[i] }

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

// Which X coordinates do these words share?
// again
//     extra
//   zit
func (w Word) sharedX(o Word) []int {
	xs := []int{}
	// Iterate through the first word, see which indexes are contained by the second
	// TODO(bsprague) There are likely more efficient interval calculation methods
	for i := 0; i < len(w.Text); i++ {
		if o.containsX(i + w.Loc.X) {
			xs = append(xs, i+w.Loc.X)
		}
	}
	return xs
}

func (w Word) containsX(i int) bool {
	// If the given index is at or after the first letter, but is before the end,
	// then it's contained
	return i >= w.Loc.X && i < (w.Loc.X+len(w.Text))
}

func (w Word) letterAt(x int) string {
	if x < w.Loc.X || x > (w.Loc.X+len(w.Text)-1) {
		return ""
	}
	i := x - w.Loc.X
	return w.Text[i : i+1]
}

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

func (b *Board) implicitWords() []string {
	// The algorithm:
	// 1. Filter down to only one orientation
	// 2. Bucket items by opposite axis (eg filter to horizontal and bucket by Y coordinate
	// 3. Look for sequential indices, build words from overlap
	// 4. Reset once overlap breaks
	h, _ := b.byOrientation()

	sort.Sort(byX(h))
	hm := make(map[int][]Word)

	for _, word := range h {
		if ws, ok := hm[word.Loc.Y]; ok {
			hm[word.Loc.Y] = append(ws, word)
		} else {
			hm[word.Loc.Y] = []Word{word}
		}
	}
	keys := make([]int, len(hm))
	i := 0
	for y := range hm {
		keys[i] = y
		i++
	}
	sort.Ints(keys)
	start := false
	//sm := make(map[int][]string)
	for i, cur := range keys[:len(keys)-1] {
		next := keys[i+1]
		// We've found words in adjacent rows
		if next-cur == 1 {
			// This is the first adjacent row
			if !start {
				//for x, s := range findOverlapX(hm[cur], hm[next]) {
				// TODO(bsprague): Fill in this section, starting new words when
				// necessary, and splitting words when necessary
				//}
			}
		}
	}

	return []string{}
}

func findOverlapX(l1, l2 []Word) map[int]string {
	m := make(map[int]string)
	for _, w1 := range l1 {
		for _, w2 := range l2 {
			if w1.Loc.X > (w2.Loc.X + len(w2.Text) - 1) {
				// Skip this one, we haven't reached the word yet
				continue
			}
			if w2.Loc.X > (w1.Loc.X + len(w1.Text) - 1) {
				// Skip this and all after, since we're sorted by X, none of these overlap
				break
			}
			for _, x := range w1.sharedX(w2) {
				// w1 and w2 overlap here, lets get their letters
				m[x] = w1.letterAt(x) + w2.letterAt(x)
			}
		}
	}
	return m
}

func (b *Board) byOrientation() (h []Word, v []Word) {
	for _, word := range b.Words {
		switch word.Orientation {
		case Horizontal:
			h = append(h, word)
		case Vertical:
			v = append(v, word)
		}
	}
	return
}
