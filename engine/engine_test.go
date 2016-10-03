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
			want: []CharLoc{{"t", Loc{1, 2}}, {"e", Loc{2, 2}}, {"s", Loc{3, 2}}, {"t", Loc{4, 2}}},
		},
		{
			word: Word{
				Orientation: Vertical,
				Text:        "again",
				Loc:         Loc{X: 2, Y: 4},
			},
			want: []CharLoc{{"a", Loc{2, 4}}, {"g", Loc{2, 5}}, {"a", Loc{2, 6}}, {"i", Loc{2, 7}}, {"n", Loc{2, 8}}},
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
	// TODO(bsprague): Test a bunch of tricky edge-cases
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
