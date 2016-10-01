package engine

type Orientation int

const (
	None Orientation = iota
	Horizontal
	Vertical
)

type Board struct {
	Words []Word
}

type Word struct {
	Orientation Orientation
	Text        string
	X, Y        uint8
}

type CharLoc struct {
	Letter string
	Loc    Loc
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
	for i := range w.Text {
		cls[i] = CharLoc{
			Letter: w.Text[i : i+1],
			Loc:    {X: w.X + i*dx, Y: w.Y + i*dy},
		}
	}
	return cls
}

type Loc struct {
	X, Y uint8
}

func (b Board) Valid(letters []string) bool {
	letterMap := make(map[Loc]string)

	for _, word := range b.Words {
		for _, cl := range word.CharLocs() {
			// If we've already placed a letter there, they've sent us a malformed
			// board
			if l, ok := letterMap; ok && l != cl.Letter {
				return false
			} else if !ok {
				letterMap[cl.Loc] = cl.Letter
			}
		}
	}
}
