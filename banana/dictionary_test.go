package banana

import (
	"strings"
	"testing"
)

func TestHasWord(t *testing.T) {
	testcases := []struct {
		input string
		want  map[string]bool
	}{
		{"cookies\ncream", map[string]bool{"cookies": true, "cream": true, "cake": false}},
		{"", map[string]bool{"cookies": false, "cream": false, "cake": false}},
		{"dogs", map[string]bool{"dog": false, "dogs": true, "doggo": false, "dogsy": false}},
	}

	for _, tc := range testcases {
		dict := NewDictionary(strings.NewReader(tc.input))
		for word, want := range tc.want {
			got := dict.HasWord(word)
			if want != got {
				t.Errorf("HasWord(%s): got %t, want %t", word, got, want)
			}
		}
	}
}
