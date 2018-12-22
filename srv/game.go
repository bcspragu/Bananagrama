package srv

import (
	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
)

var (
	wireOrientationMap = map[pb.Word_Orientation]banana.Orientation{
		pb.Word_UNKNOWN:    banana.NoOrientation,
		pb.Word_HORIZONTAL: banana.Horizontal,
		pb.Word_VERTICAL:   banana.Vertical,
	}

	engineStatusMap = map[banana.BoardStatusCode]pb.UpdateBoardResponse_Status{
		banana.Success:       pb.UpdateBoardResponse_SUCCESS,
		banana.InvalidWord:   pb.UpdateBoardResponse_INVALID_WORD,
		banana.DetachedBoard: pb.UpdateBoardResponse_DETACHED_BOARD,
		banana.NotAllLetters: pb.UpdateBoardResponse_NOT_ALL_LETTERS,
		banana.ExtraLetters:  pb.UpdateBoardResponse_EXTRA_LETTERS,
	}

	gameStatusMap = map[banana.GameStatus]pb.Game_Status{
		banana.WaitingForPlayers: pb.Game_WAITING_FOR_PLAYERS,
		banana.InProgress:        pb.Game_IN_PROGRESS,
		banana.Finished:          pb.Game_FINISHED,
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
