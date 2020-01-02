package banana

import (
	"sort"
)

// Letter represents a tile character.
type Letter rune

const (
	// noLetter is used in the absense of a valid letter to return.
	noLetter = Letter(rune(0))
)

// String turns the letter into a string.
func (l Letter) String() string {
	return string(l)
}

// Tiles represents a set of tiles.
type Tiles struct {
	freq  map[Letter]int
	count int
}

// Clone returns a deep copy of the set of tiles.
func (t *Tiles) Clone() *Tiles {
	freq := make(map[Letter]int)
	for l, idx := range t.freq {
		freq[l] = idx
	}

	return &Tiles{
		freq:  freq,
		count: t.count,
	}
}

// NewTiles returns an initialized and empty tile set.
func NewTiles() *Tiles {
	return &Tiles{
		freq: make(map[Letter]int),
	}
}

// tilesFromDistribution initializes a tile set according to a given
// distribution and scale factor.
func tilesFromDistribution(d Distribution) *Tiles {

	var (
		freq  = make(map[Letter]int)
		count = 0
	)

	for n, letters := range d {
		for _, letter := range letters {
			// Map the new index back to the letter.
			freq[letter] = n
			// Update our total count.
			count += n
		}
	}

	return &Tiles{freq: freq, count: count}
}

// Freq returns the frequency of a given letter.
func (t *Tiles) Freq(l Letter) int {
	return t.freq[l]
}

// Add adds tiles from one tile set to another, without modifying the tile set
// to add.
func (t *Tiles) Add(o *Tiles) {
	for l, freq := range o.freq {
		t.add(l, freq)
	}
}

// Inc adds one to the count of the given letter.
func (t *Tiles) Inc(l Letter) {
	t.add(l, 1)
}

// Dec removes one from the count of the given letter. This is allowed to be
// negative and is actually a key part of diffing tile sets.
func (t *Tiles) Dec(l Letter) {
	t.add(l, -1)
}

func (t *Tiles) add(l Letter, delta int) {
	t.count += delta
	t.freq[l] += delta
}

func (t *Tiles) letters() []Letter {
	ls := make([]Letter, t.count)
	i := 0
	for letter, freq := range t.freq {
		for j := 0; j < freq; j++ {
			ls[i] = letter
			i++
		}
	}
	sort.Slice(ls, func(i, j int) bool { return ls[i] < ls[j] })
	return ls
}

func (t *Tiles) AsList() []string {
	ls := make([]string, t.count)
	i := 0
	for letter, freq := range t.freq {
		str := letter.String()
		for j := 0; j < freq; j++ {
			ls[i] = str
			i++
		}
	}
	sort.Strings(ls)
	return ls
}

func (t *Tiles) Update(l Letter, freq int) {
	of := t.freq[l]
	t.freq[l] = freq
	t.count += freq - of
}

func (t *Tiles) Count() int {
	return t.count
}

func (t *Tiles) forEach(f func(l Letter, freq int) bool) {
	type letterFreq struct {
		l Letter
		f int
	}

	lfs := make([]letterFreq, len(t.freq))
	for l, freq := range t.freq {
		lfs = append(lfs, letterFreq{l: l, f: freq})
	}
	sort.Slice(lfs, func(i, j int) bool { return lfs[i].l < lfs[j].l })

	for _, lf := range lfs {
		if f(lf.l, lf.f) {
			break
		}
	}
}
