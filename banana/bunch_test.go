package banana

import (
	"math/rand"
	"testing"
)

func TestTilesFromDistribution(t *testing.T) {
	wantDistribution := map[Letter]int{
		'A': 13,
		'B': 3,
		'C': 3,
		'D': 6,
		'E': 18,
		'F': 3,
		'G': 4,
		'H': 3,
		'I': 12,
		'J': 2,
		'K': 2,
		'L': 5,
		'M': 3,
		'N': 8,
		'O': 11,
		'P': 3,
		'Q': 2,
		'R': 9,
		'S': 6,
		'T': 9,
		'U': 6,
		'V': 3,
		'W': 3,
		'X': 2,
		'Y': 3,
		'Z': 2,
	}

	tiles := tilesFromDistribution(Bananagrams())
	for letter, want := range wantDistribution {
		if got := tiles.Freq(letter); got != want {
			t.Errorf("tiles[%c]: got %d, want %d", letter, got, want)
		}
	}
}

func TestTileExhaustsAll(t *testing.T) {
	wantDistribution := map[Letter]int{
		'A': 13000,
		'B': 3000,
		'C': 3000,
		'D': 6000,
		'E': 18000,
		'F': 3000,
		'G': 4000,
		'H': 3000,
		'I': 12000,
		'J': 2000,
		'K': 2000,
		'L': 5000,
		'M': 3000,
		'N': 8000,
		'O': 11000,
		'P': 3000,
		'Q': 2000,
		'R': 9000,
		'S': 6000,
		'T': 9000,
		'U': 6000,
		'V': 3000,
		'W': 3000,
		'X': 2000,
		'Y': 3000,
		'Z': 2000,
	}
	b := newBunchWithScale(t, Bananagrams(), 1000)
	r := rand.New(rand.NewSource(0))
	// Exhaust all 144000 tiles, make sure we have none left at the end
	for i := 0; i < 144000; i++ {
		l, err := b.Tile(r)
		if err != nil {
			t.Fatalf("Tile(): %v", err)
		}
		wantDistribution[l]--
	}

	// We should have gone through all of our tiles
	want := 0
	for letter, got := range wantDistribution {
		if got != want {
			t.Errorf("%c: got %d, want %d", letter, got, want)
		}
	}
}

func TestTileNExhaustsAll(t *testing.T) {
	wantDistribution := map[Letter]int{
		'A': 13000,
		'B': 3000,
		'C': 3000,
		'D': 6000,
		'E': 18000,
		'F': 3000,
		'G': 4000,
		'H': 3000,
		'I': 12000,
		'J': 2000,
		'K': 2000,
		'L': 5000,
		'M': 3000,
		'N': 8000,
		'O': 11000,
		'P': 3000,
		'Q': 2000,
		'R': 9000,
		'S': 6000,
		'T': 9000,
		'U': 6000,
		'V': 3000,
		'W': 3000,
		'X': 2000,
		'Y': 3000,
		'Z': 2000,
	}
	b := newBunchWithScale(t, Bananagrams(), 1000)
	r := rand.New(rand.NewSource(0))
	// Exhaust all 144000 tiles, make sure we have none left at the end
	tiles, err := b.RemoveN(144000, r)
	if err != nil {
		t.Fatalf("RemoveN: %v", err)
	}
	for l, freq := range tiles.freq {
		wantDistribution[l] -= freq
	}

	// Make sure it fails when we remove another one.
	if _, err := b.RemoveN(1, r); err == nil {
		t.Error("RemoveN should have failed removing letter after exhausting all of them")
	}

	// Make sure it fails when we ask for a tile.
	l, err := b.Tile(r)
	if err == nil {
		t.Error("Tile should have failed after exhausting all of them")
	}
	if l != noLetter {
		t.Errorf("Tile = %q, want no letter", l)
	}

	// We should have gone through all of our tiles
	want := 0
	for letter, got := range wantDistribution {
		if got != want {
			t.Errorf("%c: got %d, want %d", letter, got, want)
		}
	}
}

func TestTileDistribution(t *testing.T) {
	total := 144.0
	wantDistribution := map[Letter]int{
		'A': 13,
		'B': 3,
		'C': 3,
		'D': 6,
		'E': 18,
		'F': 3,
		'G': 4,
		'H': 3,
		'I': 12,
		'J': 2,
		'K': 2,
		'L': 5,
		'M': 3,
		'N': 8,
		'O': 11,
		'P': 3,
		'Q': 2,
		'R': 9,
		'S': 6,
		'T': 9,
		'U': 6,
		'V': 3,
		'W': 3,
		'X': 2,
		'Y': 3,
		'Z': 2,
	}
	gotDistribution := map[Letter]int{
		'A': 0,
		'B': 0,
		'C': 0,
		'D': 0,
		'E': 0,
		'F': 0,
		'G': 0,
		'H': 0,
		'I': 0,
		'J': 0,
		'K': 0,
		'L': 0,
		'M': 0,
		'N': 0,
		'O': 0,
		'P': 0,
		'Q': 0,
		'R': 0,
		'S': 0,
		'T': 0,
		'U': 0,
		'V': 0,
		'W': 0,
		'X': 0,
		'Y': 0,
		'Z': 0,
	}
	b := newBunchWithScale(t, Bananagrams(), 1000)
	r := rand.New(rand.NewSource(0))
	// Get 20,000 tiles, check the distribution is what we'd expect
	for i := 0; i < 20000; i++ {
		l, err := b.Tile(r)
		if err != nil {
			t.Fatalf("Tile(): %v", err)
		}
		gotDistribution[l]++
	}

	for letter, count := range gotDistribution {
		// The percentage we're expecting, based on the actual distribution
		want := float64(wantDistribution[letter]) / total
		// Since we pulled 20,000 tiles, the percentage we got is our count / 20,000
		got := float64(count) / 20000
		diff := abs(want-got) / ((want + got) / 2)
		// If the difference is more than 10%, fail
		if diff > 0.1 {
			t.Errorf("letter %q: got %f, want %f, diff: %f%%", letter, got, want, diff*100)
		}
	}
}

func TestTileNDistribution(t *testing.T) {
	total := 144.0
	wantDistribution := map[Letter]int{
		'A': 13,
		'B': 3,
		'C': 3,
		'D': 6,
		'E': 18,
		'F': 3,
		'G': 4,
		'H': 3,
		'I': 12,
		'J': 2,
		'K': 2,
		'L': 5,
		'M': 3,
		'N': 8,
		'O': 11,
		'P': 3,
		'Q': 2,
		'R': 9,
		'S': 6,
		'T': 9,
		'U': 6,
		'V': 3,
		'W': 3,
		'X': 2,
		'Y': 3,
		'Z': 2,
	}
	gotDistribution := map[Letter]int{
		'A': 0,
		'B': 0,
		'C': 0,
		'D': 0,
		'E': 0,
		'F': 0,
		'G': 0,
		'H': 0,
		'I': 0,
		'J': 0,
		'K': 0,
		'L': 0,
		'M': 0,
		'N': 0,
		'O': 0,
		'P': 0,
		'Q': 0,
		'R': 0,
		'S': 0,
		'T': 0,
		'U': 0,
		'V': 0,
		'W': 0,
		'X': 0,
		'Y': 0,
		'Z': 0,
	}
	b := newBunchWithScale(t, Bananagrams(), 1000)
	r := rand.New(rand.NewSource(0))
	// Get 20,000 tiles, check the distribution is what we'd expect
	tiles, err := b.RemoveN(20000, r)
	if err != nil {
		t.Fatalf("RemoveN: %v", err)
	}
	for l, freq := range tiles.freq {
		gotDistribution[l] = freq
	}

	for letter, count := range gotDistribution {
		// The percentage we're expecting, based on the actual distribution
		want := float64(wantDistribution[letter]) / total
		// Since we pulled 20,000 tiles, the percentage we got is our count / 20,000
		got := float64(count) / 20000
		diff := abs(want-got) / ((want + got) / 2)
		// If the difference is more than 10%, fail
		if diff > 0.1 {
			t.Errorf("letter %c: got %f, want %f, diff: %f%%", letter, got, want, diff*100)
		}
	}
}

func TestCount(t *testing.T) {
	b := newBunchWithScale(t, Bananagrams(), 1000)
	if got, want := b.Count(), 144000; got != want {
		t.Errorf("Count: %d, want %d", got, want)
	}
}

func TestScaleErrors(t *testing.T) {
	tests := []struct {
		scale   int
		wantErr bool
	}{
		{-10, true},
		{-1, true},
		{0, true},
		{1, false},
		{10, false},
	}

	for _, test := range tests {
		_, err := Scale(Bananagrams(), test.scale)
		if test.wantErr {
			if err == nil {
				t.Error("error expected, none was returned")
			}
			continue
		}

		if err != nil {
			t.Errorf("Scale: %v", err)
		}
	}
}

func newBunchWithScale(t *testing.T, dist Distribution, scale int) *Bunch {
	scaledDist, err := Scale(dist, scale)
	if err != nil {
		t.Fatalf("Scale: %v", err)
	}

	return NewBunch(scaledDist)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
