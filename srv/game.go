package srv

import (
	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

var (
	wireOrientationMap = map[pb.Word_Orientation]banana.Orientation{
		pb.Word_UNKNOWN:    banana.None,
		pb.Word_HORIZONTAL: banana.Horizontal,
		pb.Word_VERTICAL:   banana.Vertical,
	}

	engineStatusMap = map[banana.BoardStatusCode]pb.PeelResponse_Status{
		banana.Success:       pb.PeelResponse_SUCCESS,
		banana.InvalidWord:   pb.PeelResponse_INVALID_WORD,
		banana.DetachedBoard: pb.PeelResponse_DETACHED_BOARD,
		banana.NotAllLetters: pb.PeelResponse_NOT_ALL_LETTERS,
		banana.ExtraLetters:  pb.PeelResponse_EXTRA_LETTERS,
	}
)

func boardFromWire(b *pb.Board, dict banana.Dictionary) *banana.Board {
	words := make([]banana.Word, len(b.Words))
	for i, word := range b.Words {
		words[i] = wordFromWire(word)
	}

	return &banana.Board{
		Words:      words,
		Dictionary: dict,
	}
}

func wordFromWire(w *pb.Word) banana.Word {
	return banana.Word{
		Orientation: wireOrientationMap[w.Orientation],
		Text:        w.Text,
		Loc:         banana.Loc{X: int(w.X), Y: int(w.Y)},
	}
}
