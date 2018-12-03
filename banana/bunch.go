package banana

import (
	"fmt"
	"math/rand"
)

type Bunch struct {
	tiles *Tiles
}

func (b *Bunch) Clone() *Bunch {
	return &Bunch{tiles: b.tiles.Clone()}
}

func (b *Bunch) Count() int {
	return b.tiles.count
}

func (b *Bunch) Inc(l Letter) {
	b.tiles.Inc(l)
}

// RemoveN retrieves N tiles from the bunch.
func (b *Bunch) RemoveN(n int, r *rand.Rand) (*Tiles, error) {
	if b.tiles.count < n {
		return nil, fmt.Errorf("only have %d tiles, can't remove %d tiles", b.tiles.count, n)
	}

	t := NewTiles()
	for i := 0; i < n; i++ {
		c := r.Intn(b.tiles.count)
		b.tiles.forEach(func(l Letter, freq int) bool {
			if freq > c {
				b.tiles.Dec(l)
				t.Inc(l)
				return true
			}
			c -= freq
			return false
		})
	}
	return t, nil
}

func (b *Bunch) Tile(r *rand.Rand) (Letter, error) {
	t, err := b.RemoveN(1, r)
	if err != nil {
		return noLetter, err
	}

	ls := t.letters()
	if len(ls) != 1 {
		return noLetter, fmt.Errorf("asked for tile set of size 1, got size %d", len(ls))
	}

	return ls[0], nil
}

type Distribution map[int][]Letter

// Bananagrams returns the distribution of letters on a Bananagrams board.
func Bananagrams() Distribution {
	return map[int][]Letter{
		2:  []Letter{'j', 'k', 'q', 'x', 'z'},
		3:  []Letter{'b', 'c', 'f', 'h', 'm', 'p', 'v', 'w', 'y'},
		4:  []Letter{'g'},
		5:  []Letter{'l'},
		6:  []Letter{'d', 's', 'u'},
		8:  []Letter{'n'},
		9:  []Letter{'t', 'r'},
		11: []Letter{'o'},
		12: []Letter{'i'},
		13: []Letter{'a'},
		18: []Letter{'e'},
	}
}

func NewBunch(dist Distribution, scale int) (*Bunch, error) {
	t, err := tilesFromDistribution(dist, scale)
	if err != nil {
		return nil, err
	}
	return &Bunch{tiles: t}, nil
}
