package engine

import (
	"reflect"
	"sort"
	"testing"
)

func TestCharLocs(t *testing.T) {
	testcases := []struct {
		word Word
		want []CharLoc
	}{
		{
			word: Word{
				Orientation: Horizontal,
				Text:        "test",
				Loc:         Loc{X: 1, Y: 2},
			},
			want: []CharLoc{{'t', Loc{1, 2}}, {'e', Loc{2, 2}}, {'s', Loc{3, 2}}, {'t', Loc{4, 2}}},
		},
		{
			word: Word{
				Orientation: Vertical,
				Text:        "again",
				Loc:         Loc{X: 2, Y: 4},
			},
			want: []CharLoc{{'a', Loc{2, 4}}, {'g', Loc{2, 5}}, {'a', Loc{2, 6}}, {'i', Loc{2, 7}}, {'n', Loc{2, 8}}},
		},
	}

	for _, tc := range testcases {
		got := tc.word.CharLocs()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("word %s: got %v, want %v", tc.word.Text, got, tc.want)
		}
	}
}

func TestFindWords(t *testing.T) {
	testcases := []struct {
		words []Word
		want  []string
	}{
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}},
			want:  []string{"cat"},
		},
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}, {Horizontal, "bad", Loc{1, 1}}, {Horizontal, "smell", Loc{1, 2}}},
			want:  []string{"cat", "abs", "tam", "bad", "smell", "de"},
		},
		{
			// Inverting the order of the words shows that we're sorting beforehand
			words: []Word{{Vertical, "radish", Loc{0, 5}}, {Vertical, "horse", Loc{0, 0}}},
			want:  []string{"horseradish"},
		},
		{
			words: []Word{{Horizontal, "h", Loc{0, 0}}, {Vertical, "i", Loc{0, 1}}},
			want:  []string{"hi"},
		},
		{
			words: []Word{{Horizontal, "hello", Loc{3, 0}}, {Vertical, "hello", Loc{0, 3}}},
			want:  []string{"hello", "hello"},
		},
		{
			words: []Word{{Vertical, "hello", Loc{3, 0}}, {Horizontal, "hello", Loc{0, 3}}},
			want:  []string{"hello", "hello"},
		},
		{
			words: []Word{{Vertical, "muffin", Loc{0, 0}}, {Horizontal, "scent", Loc{0, 6}}},
			want:  []string{"muffins", "scent"},
		},
	}

	for _, tc := range testcases {
		got := (&Board{Words: tc.words}).findWords()
		sort.Strings(got)
		sort.Strings(tc.want)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestConnected(t *testing.T) {
	testcases := []struct {
		words []Word
		want  bool
	}{
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}},
			want:  true,
		},
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}, {Horizontal, "bad", Loc{1, 1}}, {Horizontal, "smell", Loc{1, 2}}},
			want:  true,
		},
		{
			// horse radish, note the space
			words: []Word{{Vertical, "radish", Loc{0, 6}}, {Vertical, "horse", Loc{0, 0}}},
			want:  false,
		},
		{
			words: []Word{{Horizontal, "hello", Loc{0, 0}}, {Vertical, "eggs", Loc{1, 0}}, {Horizontal, "woof", Loc{1, 6}}, {Vertical, "wood", Loc{1, 6}}},
			want:  false,
		},
	}

	for _, tc := range testcases {
		board := &Board{Words: tc.words}
		if !board.precompute() {
			t.Errorf("couldn't precompute board %v", board)
		}
		got := board.connected()
		if got != tc.want {
			t.Errorf("connectd(%v): got %v, want %v", board, got, tc.want)
		}
	}
}

func TestContainsExactly(t *testing.T) {
	testcases := []struct {
		words   []Word
		letters map[rune]int
		want    bool
	}{
		{
			words:   []Word{{Horizontal, "cat", Loc{0, 0}}},
			letters: map[rune]int{'c': 1, 'a': 1, 't': 1},
			want:    true,
		},
		{
			words:   []Word{{Horizontal, "hello", Loc{0, 0}}, {Horizontal, "hello", Loc{0, 0}}},
			letters: map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
			want:    true,
		},
		{
			words:   []Word{{Horizontal, "hello", Loc{0, 0}}, {Horizontal, "hello", Loc{0, 0}}},
			letters: map[rune]int{'h': 2, 'e': 2, 'l': 4, 'o': 2},
			want:    false,
		},
		{
			words:   []Word{},
			letters: map[rune]int{'h': 2, 'e': 2, 'l': 4, 'o': 2},
			want:    false,
		},
		{
			words:   []Word{{Vertical, "hello", Loc{3, 0}}, {Horizontal, "hello", Loc{0, 3}}},
			letters: map[rune]int{'h': 2, 'e': 2, 'l': 3, 'o': 2},
			want:    true,
		},
		{
			words:   []Word{{Vertical, "dodge", Loc{0, 0}}, {Horizontal, "douse", Loc{0, 0}}, {Horizontal, "deemed", Loc{0, 2}}, {Horizontal, "each", Loc{0, 4}}},
			letters: map[rune]int{'a': 1, 'c': 1, 'd': 3, 'e': 5, 'g': 1, 'h': 1, 'm': 1, 'o': 2, 's': 1, 'u': 1},
			want:    true,
		},
		{
			words:   []Word{{Vertical, "zalgo", Loc{0, 0}}, {Vertical, "quick", Loc{1, 0}}, {Vertical, "brave", Loc{2, 0}}},
			letters: map[rune]int{'a': 2, 'b': 1, 'c': 1, 'e': 1, 'g': 1, 'i': 1, 'k': 1, 'l': 1, 'o': 1, 'q': 1, 'r': 1, 'u': 1, 'v': 1, 'z': 1},
			want:    true,
		},
	}

	for _, tc := range testcases {
		board := &Board{Words: tc.words}
		if !board.precompute() {
			t.Errorf("couldn't precompute board %v", board)
		}
		got := board.containsExactly(freqListFromMap(tc.letters))
		if got != tc.want {
			t.Errorf("containsExactly(%v): got %v, want %v", board, got, tc.want)
		}
	}
}

func freqListFromMap(m map[rune]int) FreqList {
	var f FreqList
	for l, c := range m {
		// Don't try this at home kids
		f[int(l)-97] = c
	}
	return f
}
