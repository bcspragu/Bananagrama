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
	orientationMap = map[banana.Orientation]pb.Word_Orientation{
		banana.NoOrientation: pb.Word_UNKNOWN,
		banana.Horizontal:    pb.Word_HORIZONTAL,
		banana.Vertical:      pb.Word_VERTICAL,
	}

	gameStatusMap = map[banana.GameStatus]pb.GameStatus{
		banana.WaitingForPlayers: pb.GameStatus_WAITING_FOR_PLAYERS,
		banana.InProgress:        pb.GameStatus_IN_PROGRESS,
		banana.Finished:          pb.GameStatus_FINISHED,
	}
)

func boardToWire(b *banana.Board) *pb.Board {
	words := b.Words()
	pbWords := make([]*pb.Word, len(words))
	for i, word := range words {
		pbWords[i] = wordToWire(word)
	}

	return &pb.Board{
		Words: pbWords,
	}
}

func wordToWire(w banana.Word) *pb.Word {
	return &pb.Word{
		Orientation: orientationMap[w.Orientation],
		Text:        w.Text,
		X:           int32(w.Loc.X),
		Y:           int32(w.Loc.Y),
	}
}

func boardFromWire(c *banana.Config, pbWords []*pb.Word) (*banana.Board, error) {
	words := make([]banana.Word, len(pbWords))
	for i, pbWord := range pbWords {
		words[i] = wordFromWire(pbWord)
	}

	return banana.NewBoardWithWords(c, words)
}

func wordFromWire(w *pb.Word) banana.Word {
	return banana.Word{
		Orientation: wireOrientationMap[w.Orientation],
		Text:        w.Text,
		Loc:         banana.Loc{X: int(w.X), Y: int(w.Y)},
	}
}

func charLocsListToWire(cls []banana.CharLocs) []*pb.CharLocs {
	wireCls := make([]*pb.CharLocs, len(cls))
	for i, cl := range cls {
		wireCls[i] = charLocsToWire(cl)
	}

	return wireCls
}

func charLocsToWire(cls banana.CharLocs) *pb.CharLocs {
	locs := make([]*pb.CharLoc, len(cls.Locs))
	for i, cl := range cls.Locs {
		locs[i] = charLocToWire(cl)
	}

	return &pb.CharLocs{
		Text: cls.Word,
		Locs: locs,
	}
}

func charLocToWire(cl banana.CharLoc) *pb.CharLoc {
	return &pb.CharLoc{
		Letter: string(cl.Letter),
		X:      int32(cl.Loc.X),
		Y:      int32(cl.Loc.Y),
	}
}
