package pubsub

import "github.com/bcspragu/Bananagrama/banana"

type PayloadType int

const (
	UnknownPayload PayloadType = iota
	// A game has been created or updated.
	PayloadTypeGameUpdated
	// The game has started.
	PayloadTypeGameStart
	// The game has ended.
	PayloadTypeGameOver
	// A player has played a word.
	PayloadTypePlayerMove
	// A player has dumped.
	PayloadTypePlayerDump
	// A player has played an invalid word.
	PayloadTypeWordFail
	// A player's board is in a rough state.
	PayloadTypeBoardFail
)

type Payload struct {
	Type PayloadType
	// The below messages are only populated for their given
	// type.
	GameUpdated *GameUpdated
	GameOver    *GameOver
	PlayerMove  *PlayerMove
	PlayerDump  *PlayerDump
	WordFail    *WordFail
	BoardFail   *BoardFail
}

type GameUpdated struct {
	ID          banana.GameID
	Name        string
	PlayerCount int
	Status      banana.GameStatus
}

type GameOver struct {
	Winner string
}

type PlayerMove struct {
	ID   banana.PlayerID
	Name string
	Word string
}

type PlayerDump struct {
	Name string
}

type WordFail struct {
	Name string
	Word string
}

type BoardFail struct {
	Name string
}
