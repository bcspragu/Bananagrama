package srv

import (
	"reflect"
	"testing"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

func TestWordFromWire(t *testing.T) {
	testcases := []struct {
		in   *pb.Word
		want banana.Word
	}{
		{
			in: &pb.Word{
				Text:        "test",
				Orientation: pb.Word_UNKNOWN,
				X:           1,
				Y:           2,
			},
			want: banana.Word{
				Text:        "test",
				Orientation: banana.NoOrientation,
				Loc:         banana.Loc{1, 2},
			},
		},
		{
			in: &pb.Word{
				Text:        "fudge",
				Orientation: pb.Word_HORIZONTAL,
				X:           0,
				Y:           0,
			},
			want: banana.Word{
				Text:        "fudge",
				Orientation: banana.Horizontal,
				Loc:         banana.Loc{0, 0},
			},
		},
		{
			in: &pb.Word{
				Text:        "globble",
				Orientation: pb.Word_VERTICAL,
				X:           23,
				Y:           2,
			},
			want: banana.Word{
				Text:        "globble",
				Orientation: banana.Vertical,
				Loc:         banana.Loc{23, 2},
			},
		},
	}

	for _, tc := range testcases {
		if got := wordFromWire(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("wordFromWire(%v): %v, want %v", tc.in, got, tc.want)
		}
	}
}
