package banana

import (
	"reflect"
	"testing"
)

func TestAsList(t *testing.T) {
	testcases := []struct {
		letters map[Letter]int
		want    []string
	}{
		{
			letters: map[Letter]int{'c': 1, 'a': 1, 't': 1},
			want:    []string{"a", "c", "t"},
		},
		{
			letters: map[Letter]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
			want:    []string{"e", "h", "l", "l", "o"},
		},
		{
			letters: map[Letter]int{'h': 2, 'e': 2, 'l': 4, 'o': 2},
			want:    []string{"e", "e", "h", "h", "l", "l", "l", "l", "o", "o"},
		},
	}

	for _, tc := range testcases {
		fl := tilesFromMap(tc.letters)
		got := fl.AsList()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AsList(%v): got %v, want %v", tc.letters, got, tc.want)
		}
	}
}

func TestAdd(t *testing.T) {
	t1 := tilesFromMap(map[Letter]int{'c': 1, 'a': 1, 't': 1})
	t2 := tilesFromMap(map[Letter]int{'a': 1, 't': 2})
	want := map[Letter]int{'c': 1, 'a': 2, 't': 3}

	t1.Add(t2)

	for l, freq := range t1.freq {
		want[l] -= freq
	}

	for l, freq := range want {
		if freq < 0 {
			t.Errorf("letter %q was represented %d too many times", l, -freq)
		}
		if freq > 0 {
			t.Errorf("letter %q was represented %d too few times", l, freq)
		}
	}
}

func tilesFromMap(m map[Letter]int) *Tiles {
	t := newTiles()
	for l, c := range m {
		t.Update(Letter(l), c)
	}
	return t
}
