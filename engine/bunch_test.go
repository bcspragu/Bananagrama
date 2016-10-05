package engine

import (
	"math/rand"
	"testing"
)

func TestNewBunch(t *testing.T) {
	wantDistribution := map[Letter]int{
		'a': 13000,
		'b': 3000,
		'c': 3000,
		'd': 6000,
		'e': 18000,
		'f': 3000,
		'g': 4000,
		'h': 3000,
		'i': 12000,
		'j': 2000,
		'k': 2000,
		'l': 5000,
		'm': 3000,
		'n': 8000,
		'o': 11000,
		'p': 3000,
		'q': 2000,
		'r': 9000,
		's': 6000,
		't': 9000,
		'u': 6000,
		'v': 3000,
		'w': 3000,
		'x': 2000,
		'y': 3000,
		'z': 2000,
	}
	b := NewBunch()
	for letter, want := range wantDistribution {
		got := b.tiles.Freq(letter)
		if got != want {
			t.Errorf("tiles[%c]: got %d, want %d", letter, got, want)
		}
	}
}

func TestTileExhaustsAll(t *testing.T) {
	wantDistribution := map[Letter]int{
		'a': 13000,
		'b': 3000,
		'c': 3000,
		'd': 6000,
		'e': 18000,
		'f': 3000,
		'g': 4000,
		'h': 3000,
		'i': 12000,
		'j': 2000,
		'k': 2000,
		'l': 5000,
		'm': 3000,
		'n': 8000,
		'o': 11000,
		'p': 3000,
		'q': 2000,
		'r': 9000,
		's': 6000,
		't': 9000,
		'u': 6000,
		'v': 3000,
		'w': 3000,
		'x': 2000,
		'y': 3000,
		'z': 2000,
	}
	b := NewBunch()
	b.rand = rand.New(rand.NewSource(0))
	// Exhaust all 144000 tiles, make sure we have none left at the end
	for i := 0; i < 144000; i++ {
		wantDistribution[b.Tile()]--
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
		'a': 13,
		'b': 3,
		'c': 3,
		'd': 6,
		'e': 18,
		'f': 3,
		'g': 4,
		'h': 3,
		'i': 12,
		'j': 2,
		'k': 2,
		'l': 5,
		'm': 3,
		'n': 8,
		'o': 11,
		'p': 3,
		'q': 2,
		'r': 9,
		's': 6,
		't': 9,
		'u': 6,
		'v': 3,
		'w': 3,
		'x': 2,
		'y': 3,
		'z': 2,
	}
	gotDistribution := map[Letter]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}
	b := NewBunch()
	b.rand = rand.New(rand.NewSource(0))
	// Get 20,000 tiles, check the distribution is what we'd expect
	for i := 0; i < 20000; i++ {
		gotDistribution[b.Tile()]++
	}

	for letter, count := range gotDistribution {
		// The percentage we're expecting, based on the actual distribution
		want := float64(wantDistribution[letter]) / total
		// Since we pulled 20,000 tiles, the percentage we got is our count / 20,000
		got := float64(count) / 20000
		diff := abs(want-got) / ((want + got) / 2)
		// If the difference is more than 10%, fail
		if diff > 0.1 {
			t.Errorf("letter %c: got %f, want %f, diff: %f%", letter, got, want, diff*100)
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
