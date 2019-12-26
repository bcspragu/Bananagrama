package banana

import (
	"reflect"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		want  []CharLocs
	}{
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}},
			want: []CharLocs{
				CharLocs{
					Word: "cat",
					Locs: []CharLoc{
						{Letter: 'c'},
						{Letter: 'a', Loc: Loc{X: 1}},
						{Letter: 't', Loc: Loc{X: 2}},
					},
				},
			},
		},
		{
			words: []Word{{Horizontal, "cat", Loc{0, 0}}, {Horizontal, "bad", Loc{1, 1}}, {Horizontal, "smell", Loc{1, 2}}},
			want: []CharLocs{
				CharLocs{Word: "abs", Locs: []CharLoc{
					{Letter: 'a', Loc: Loc{X: 1}},
					{Letter: 'b', Loc: Loc{X: 1, Y: 1}},
					{Letter: 's', Loc: Loc{X: 1, Y: 2}},
				},
				},
				CharLocs{Word: "bad", Locs: []CharLoc{
					{Letter: 'b', Loc: Loc{X: 1, Y: 1}},
					{Letter: 'a', Loc: Loc{X: 2, Y: 1}},
					{Letter: 'd', Loc: Loc{X: 3, Y: 1}},
				},
				},
				CharLocs{Word: "cat", Locs: []CharLoc{
					{Letter: 'c'},
					{Letter: 'a', Loc: Loc{X: 1}},
					{Letter: 't', Loc: Loc{X: 2}},
				},
				},
				CharLocs{Word: "de", Locs: []CharLoc{
					{Letter: 'd', Loc: Loc{X: 3, Y: 1}},
					{Letter: 'e', Loc: Loc{X: 3, Y: 2}},
				},
				},
				CharLocs{Word: "smell", Locs: []CharLoc{
					{Letter: 's', Loc: Loc{X: 1, Y: 2}},
					{Letter: 'm', Loc: Loc{X: 2, Y: 2}},
					{Letter: 'e', Loc: Loc{X: 3, Y: 2}},
					{Letter: 'l', Loc: Loc{X: 4, Y: 2}},
					{Letter: 'l', Loc: Loc{X: 5, Y: 2}},
				},
				},
				CharLocs{Word: "tam", Locs: []CharLoc{
					{Letter: 't', Loc: Loc{X: 2}},
					{Letter: 'a', Loc: Loc{X: 2, Y: 1}},
					{Letter: 'm', Loc: Loc{X: 2, Y: 2}},
				},
				},
			},
		},
		{
			words: []Word{{Vertical, "cat", Loc{0, 0}}, {Vertical, "bad", Loc{1, 1}}, {Vertical, "smell", Loc{2, 1}}},
			want: []CharLocs{
				CharLocs{
					Word: "abs",
					Locs: []CharLoc{{Letter: 'a', Loc: Loc{Y: 1}}, {Letter: 'b', Loc: Loc{X: 1, Y: 1}}, {Letter: 's', Loc: Loc{X: 2, Y: 1}}},
				},
				CharLocs{
					Word: "bad",
					Locs: []CharLoc{{Letter: 'b', Loc: Loc{X: 1, Y: 1}}, {Letter: 'a', Loc: Loc{X: 1, Y: 2}}, {Letter: 'd', Loc: Loc{X: 1, Y: 3}}},
				},
				CharLocs{
					Word: "cat",
					Locs: []CharLoc{{Letter: 'c'}, {Letter: 'a', Loc: Loc{Y: 1}}, {Letter: 't', Loc: Loc{Y: 2}}},
				},
				CharLocs{
					Word: "de",
					Locs: []CharLoc{{Letter: 'd', Loc: Loc{X: 1, Y: 3}}, {Letter: 'e', Loc: Loc{X: 2, Y: 3}}},
				},
				CharLocs{
					Word: "smell",
					Locs: []CharLoc{{Letter: 's', Loc: Loc{X: 2, Y: 1}}, {Letter: 'm', Loc: Loc{X: 2, Y: 2}}, {Letter: 'e', Loc: Loc{X: 2, Y: 3}}, {Letter: 'l', Loc: Loc{X: 2, Y: 4}}, {Letter: 'l', Loc: Loc{X: 2, Y: 5}}},
				},
				CharLocs{
					Word: "tam",
					Locs: []CharLoc{{Letter: 't', Loc: Loc{Y: 2}}, {Letter: 'a', Loc: Loc{X: 1, Y: 2}}, {Letter: 'm', Loc: Loc{X: 2, Y: 2}}},
				},
			},
		},
		{
			// Inverting the order of the words shows that we're sorting beforehand
			words: []Word{{Vertical, "radish", Loc{0, 5}}, {Vertical, "horse", Loc{0, 0}}},
			want: []CharLocs{
				CharLocs{
					Word: "horseradish",
					Locs: []CharLoc{{Letter: 'h'}, {Letter: 'o', Loc: Loc{Y: 1}}, {Letter: 'r', Loc: Loc{Y: 2}}, {Letter: 's', Loc: Loc{Y: 3}}, {Letter: 'e', Loc: Loc{Y: 4}}, {Letter: 'r', Loc: Loc{Y: 5}}, {Letter: 'a', Loc: Loc{Y: 6}}, {Letter: 'd', Loc: Loc{Y: 7}}, {Letter: 'i', Loc: Loc{Y: 8}}, {Letter: 's', Loc: Loc{Y: 9}}, {Letter: 'h', Loc: Loc{Y: 10}}}},
			},
		},
		{
			words: []Word{{Horizontal, "h", Loc{0, 0}}, {Vertical, "i", Loc{0, 1}}},
			want: []CharLocs{
				CharLocs{Word: "hi", Locs: []CharLoc{{Letter: 'h'}, {Letter: 'i', Loc: Loc{Y: 1}}}},
			},
		},
		{
			words: []Word{{Horizontal, "hello", Loc{3, 0}}, {Vertical, "hello", Loc{0, 3}}},
			want: []CharLocs{
				CharLocs{Word: "hello", Locs: []CharLoc{{Letter: 'h', Loc: Loc{Y: 3}}, {Letter: 'e', Loc: Loc{Y: 4}}, {Letter: 'l', Loc: Loc{Y: 5}}, {Letter: 'l', Loc: Loc{Y: 6}}, {Letter: 'o', Loc: Loc{Y: 7}}}},
				CharLocs{Word: "hello", Locs: []CharLoc{{Letter: 'h', Loc: Loc{X: 3}}, {Letter: 'e', Loc: Loc{X: 4}}, {Letter: 'l', Loc: Loc{X: 5}}, {Letter: 'l', Loc: Loc{X: 6}}, {Letter: 'o', Loc: Loc{X: 7}}}},
			},
		},
		{
			words: []Word{{Vertical, "hello", Loc{3, 0}}, {Horizontal, "hello", Loc{0, 3}}},
			want: []CharLocs{
				CharLocs{Word: "hello", Locs: []CharLoc{{Letter: 'h', Loc: Loc{X: 3}}, {Letter: 'e', Loc: Loc{X: 3, Y: 1}}, {Letter: 'l', Loc: Loc{X: 3, Y: 2}}, {Letter: 'l', Loc: Loc{X: 3, Y: 3}}, {Letter: 'o', Loc: Loc{X: 3, Y: 4}}}},
				CharLocs{Word: "hello", Locs: []CharLoc{{Letter: 'h', Loc: Loc{Y: 3}}, {Letter: 'e', Loc: Loc{X: 1, Y: 3}}, {Letter: 'l', Loc: Loc{X: 2, Y: 3}}, {Letter: 'l', Loc: Loc{X: 3, Y: 3}}, {Letter: 'o', Loc: Loc{X: 4, Y: 3}}}},
			},
		},
		{
			words: []Word{{Vertical, "muffin", Loc{0, 0}}, {Horizontal, "scent", Loc{0, 6}}},
			want: []CharLocs{
				CharLocs{Word: "muffins", Locs: []CharLoc{{Letter: 'm'}, {Letter: 'u', Loc: Loc{Y: 1}}, {Letter: 'f', Loc: Loc{Y: 2}}, {Letter: 'f', Loc: Loc{Y: 3}}, {Letter: 'i', Loc: Loc{Y: 4}}, {Letter: 'n', Loc: Loc{Y: 5}}, {Letter: 's', Loc: Loc{Y: 6}}}},
				CharLocs{Word: "scent", Locs: []CharLoc{{Letter: 's', Loc: Loc{Y: 6}}, {Letter: 'c', Loc: Loc{X: 1, Y: 6}}, {Letter: 'e', Loc: Loc{X: 2, Y: 6}}, {Letter: 'n', Loc: Loc{X: 3, Y: 6}}, {Letter: 't', Loc: Loc{X: 4, Y: 6}}}},
			},
		},
		{
			words: []Word{
				{Horizontal, "g  l", Loc{0, 0}},
				{Horizontal, "oa", Loc{1, 0}},
			},
			want: []CharLocs{
				CharLocs{Word: "goal", Locs: []CharLoc{{Letter: 'g'}, {Letter: 'o', Loc: Loc{X: 1}}, {Letter: 'a', Loc: Loc{X: 2}}, {Letter: 'l', Loc: Loc{X: 3}}}},
			},
		},
	}

	for _, tc := range testcases {
		got := (&Board{Words: tc.words}).findWords()

		sort.Slice(got, func(i, j int) bool { return got[i].Word < got[j].Word })
		sort.Slice(tc.want, func(i, j int) bool { return tc.want[i].Word < tc.want[j].Word })

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("findWords (-want +got):\n%s", diff)
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
			// Diagonal is not connected
			words: []Word{{Horizontal, "cat", Loc{0, 0}}, {Horizontal, "bad", Loc{3, 1}}},
			want:  false,
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
		if err := board.precompute(); err != nil {
			t.Errorf("couldn't precompute board %v: %v", board, err)
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
		letters map[Letter]int
		want    bool
	}{
		{
			words:   []Word{{Horizontal, "cat", Loc{0, 0}}},
			letters: map[Letter]int{'c': 1, 'a': 1, 't': 1},
			want:    true,
		},
		{
			words:   []Word{{Horizontal, "hello", Loc{0, 0}}, {Horizontal, "hello", Loc{0, 0}}},
			letters: map[Letter]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
			want:    true,
		},
		{
			words:   []Word{{Horizontal, "hello", Loc{0, 0}}, {Horizontal, "hello", Loc{0, 0}}},
			letters: map[Letter]int{'h': 2, 'e': 2, 'l': 4, 'o': 2},
			want:    false,
		},
		{
			words:   []Word{},
			letters: map[Letter]int{'h': 2, 'e': 2, 'l': 4, 'o': 2},
			want:    false,
		},
		{
			words:   []Word{{Vertical, "hello", Loc{3, 0}}, {Horizontal, "hello", Loc{0, 3}}},
			letters: map[Letter]int{'h': 2, 'e': 2, 'l': 3, 'o': 2},
			want:    true,
		},
		{
			words:   []Word{{Vertical, "dodge", Loc{0, 0}}, {Horizontal, "douse", Loc{0, 0}}, {Horizontal, "deemed", Loc{0, 2}}, {Horizontal, "each", Loc{0, 4}}},
			letters: map[Letter]int{'a': 1, 'c': 1, 'd': 3, 'e': 5, 'g': 1, 'h': 1, 'm': 1, 'o': 2, 's': 1, 'u': 1},
			want:    true,
		},
		{
			words:   []Word{{Vertical, "zalgo", Loc{0, 0}}, {Vertical, "quick", Loc{1, 0}}, {Vertical, "brave", Loc{2, 0}}},
			letters: map[Letter]int{'a': 2, 'b': 1, 'c': 1, 'e': 1, 'g': 1, 'i': 1, 'k': 1, 'l': 1, 'o': 1, 'q': 1, 'r': 1, 'u': 1, 'v': 1, 'z': 1},
			want:    true,
		},
	}

	for _, tc := range testcases {
		board := &Board{Words: tc.words}
		if err := board.precompute(); err != nil {
			t.Errorf("couldn't precompute board %v: %v", board, err)
		}
		got := board.containsExactly(tilesFromMap(tc.letters))
		if got != tc.want {
			t.Errorf("containsExactly(%v): got %v, want %v", board, got, tc.want)
		}
	}
}

func TestMalformedPrecompute(t *testing.T) {
	board := &Board{
		Words: []Word{
			{Horizontal, "test", Loc{0, 0}},
			{Vertical, "bad", Loc{1, 0}},
		},
	}

	if err := board.precompute(); err == nil {
		t.Error("expected precompute to fail")
	}
}

func TestValidate(t *testing.T) {
	dict := &dictImpl{
		words: map[string]struct{}{
			"bats":    struct{}{},
			"cats":    struct{}{},
			"catsit":  struct{}{},
			"attack":  struct{}{},
			"attacks": struct{}{},
			"goal":    struct{}{},
		},
	}
	tests := []struct {
		desc    string
		tiles   *Tiles
		board   *Board
		want    *BoardValidation
		wantErr bool
	}{
		{
			desc: "success",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 1,
					'k': 1,
					's': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attack", Loc{0, 0}},
					{Vertical, "cats", Loc{4, 0}},
				},
			},
			want: &BoardValidation{},
		},
		{
			desc: "invalid board",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 1,
					'b': 1,
					'k': 1,
					's': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attack", Loc{0, 0}},
					{Vertical, "bats", Loc{4, 0}},
				},
			},
			wantErr: true,
		},
		{
			desc: "not all letters used",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 1,
					'k': 1,
					's': 4,
					'q': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attacks", Loc{0, 0}},
					{Vertical, "cats", Loc{4, 0}},
				},
			},
			want: &BoardValidation{
				UnusedLetters: []string{"q", "s"},
			},
		},
		{
			desc: "used letters they didn't have",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 1,
					'k': 1,
					's': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attacks", Loc{0, 0}},
					{Vertical, "catsit", Loc{4, 0}},
				},
			},
			want: &BoardValidation{
				ExtraLetters: []string{"i", "s", "t"},
			},
		},
		{
			desc: "not in dictionary",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 1,
					'k': 1,
					'e': 1,
					'd': 1,
					's': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attacked", Loc{0, 0}},
					{Vertical, "cats", Loc{4, 0}},
				},
			},
			want: &BoardValidation{
				InvalidWords: []CharLocs{
					CharLocs{Word: "attacked", Locs: []CharLoc{{Letter: 'a'}, {Letter: 't', Loc: Loc{X: 1}}, {Letter: 't', Loc: Loc{X: 2}}, {Letter: 'a', Loc: Loc{X: 3}}, {Letter: 'c', Loc: Loc{X: 4}}, {Letter: 'k', Loc: Loc{X: 5}}, {Letter: 'e', Loc: Loc{X: 6}}, {Letter: 'd', Loc: Loc{X: 7}}}},
				},
			},
		},
		{
			desc: "not connected",
			tiles: &Tiles{
				freq: map[Letter]int{
					'a': 3,
					't': 3,
					'c': 2,
					'k': 1,
					's': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "attack", Loc{0, 0}},
					{Vertical, "cats", Loc{4, 2}},
				},
			},
			want: &BoardValidation{DetachedBoard: true},
		},
		{
			desc: "test sub-word bug",
			tiles: &Tiles{
				freq: map[Letter]int{
					'g': 1,
					'o': 1,
					'a': 1,
					'l': 1,
				},
				count: 6,
			},
			board: &Board{
				Words: []Word{
					{Horizontal, "g  l", Loc{0, 0}},
					{Horizontal, "oa", Loc{1, 0}},
				},
			},
			want: &BoardValidation{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := test.board.Validate(test.tiles, dict)
			if err != nil {
				if test.wantErr {
					// This is what we wanted.
					return
				}
				t.Errorf("ValidateBoard: %v", err)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("ValidateBoard (-want +got):\n%s", diff)
			}
		})
	}
}

func TestStartingTileCount(t *testing.T) {
	tests := []struct {
		players int
		scale   int
		want    int
	}{
		{players: 1, scale: 1, want: 21},
		{players: 2, scale: 10, want: 210},
		{players: 3, scale: 100, want: 2100},
		{players: 4, scale: 1000, want: 21000},
		{players: 5, scale: 1, want: 15},
		{players: 6, scale: 10, want: 150},
		{players: 7, scale: 1, want: 11},
		{players: 8, scale: 10, want: 110},
		{players: 9, scale: 1, want: 11},
	}

	for _, test := range tests {
		if got := StartingTileCount(test.players, test.scale); got != test.want {
			t.Errorf("StartingTileCount(%d, %d) = %d, want %d)", test.players, test.scale, got, test.want)
		}
	}
}
