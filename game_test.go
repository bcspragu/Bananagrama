package main

import (
	"reflect"
	"testing"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
	capnp "zombiezen.com/go/capnproto2"
)

func TestWordFromWire(t *testing.T) {
	testcases := []struct {
		wire potassium.Word
		want engine.Word
	}{
		{makeWireWord("test", potassium.Word_Orientation_unknown, 1, 2, t), engine.Word{engine.None, "test", engine.Loc{1, 2}}},
		{makeWireWord("fudge", potassium.Word_Orientation_horizontal, 0, 0, t), engine.Word{engine.Horizontal, "fudge", engine.Loc{0, 0}}},
		{makeWireWord("globble", potassium.Word_Orientation_vertical, 23, 2, t), engine.Word{engine.Vertical, "globble", engine.Loc{23, 2}}},
	}

	for _, tc := range testcases {
		got, err := wordFromWire(tc.wire)
		if err != nil {
			t.Errorf("wordFromWire(%v): %v", tc.wire, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("wordFromWire(%v): %v, want %v", tc.wire, got, tc.want)
		}
	}
}

func makeWireWord(text string, orientation potassium.Word_Orientation, x, y int, t *testing.T) potassium.Word {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		t.Errorf("makeWireWord(%s, %s, %d, %d): %v", text, orientation, x, y, err)
	}

	word, err := potassium.NewWord(seg)
	if err != nil {
		t.Errorf("makeWireWord(%s, %s, %d, %d): %v", text, orientation, x, y, err)
	}

	err = word.SetText(text)
	if err != nil {
		t.Errorf("makeWireWord(%s, %s, %d, %d): %v", text, orientation, x, y, err)
	}

	word.SetOrientation(orientation)
	word.SetX(uint32(x))
	word.SetY(uint32(y))

	return word
}
