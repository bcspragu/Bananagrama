package pubsub

import "github.com/bcspragu/Bananagrama/banana"

type PayloadType int

const (
	UnknownPayload PayloadType = iota
	// A game has been created or a player has been added to it.
	PayloadTypeGameUpdated
	// The game has started.
	PayloadTypeGameStarted
	// The game has ended.
	PayloadTypeGameEnded
	// A player has joined a game.
	PayloadTypePlayerJoined
	// A player has played a word.
	PayloadTypePlayerMove
	// A player has dumped.
	PayloadTypePlayerDump
	// A player has played an invalid word.
	PayloadTypeWordFail
	// A player's board is in a rough state.
	PayloadTypeBoardFail
	// The set of this player's tiles.
	PayloadTypeSelfTileUpdate
	// The state of another player's tiles.
	PayloadTypeOtherTileUpdates
)

type Payload struct {
	Type PayloadType
	// The below messages are only populated for their given
	// type.
	GameUpdated      *GameUpdated
	GameStarted      *GameStarted
	GameEnded        *GameEnded
	PlayerJoined     *PlayerJoined
	PlayerMove       *PlayerMove
	PlayerDump       *PlayerDump
	SelfTileUpdate   *SelfTileUpdate
	OtherTileUpdates *OtherTileUpdates
}

type GameUpdated struct {
	ID          banana.GameID
	Name        string
	Status      banana.GameStatus
	PlayerCount int
}

type GameStarted struct {
	NumStartingTiles int
}

type GameEnded struct {
	Winner string
}

type PlayerJoined struct {
	ID   banana.PlayerID
	Name string
}

type PlayerMove struct {
	ID           banana.PlayerID
	Name         string
	Word         string
	WordValid    bool
	BoardValid   bool
	TilesInBunch int
}

type PlayerDump struct {
	ID   banana.PlayerID
	Name string
}

type SelfTileUpdate struct {
	Tiles    *banana.Tiles
	PeelFrom banana.PlayerID
}

type OtherTileUpdates struct {
	Updates        []*OtherTileUpdate
	RemainingTiles int
}

type OtherTileUpdate struct {
	ID           banana.PlayerID
	TilesInHand  int
	TilesInBoard int
}
