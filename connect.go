package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"sort"
	"sync"
	"time"

	capnp "zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/rpc"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
)

// baseAction is the base of every message we send to the client over
// WebSockets
type baseAction struct {
	Action string `json:"action"`
}

type playersAction struct {
	baseAction
	HTML string `json:"html"`
}

const (
	// How many people can join in a single game
	MaxPlayers = 8
	// How many tiles a player receives on a dump
	DumpSize = 3
	// How many tiles each player receives on a peel. Note, there are some
	// interesting implications to how this is set. Excluding dumps, the game
	// will end after roughly
	// 144*TileScalingFactor/(PeelSize*NumConnectedPlayers) peels, and the number
	// of peels determines other things, like size of database, type of game
	// play, and amount of network traffic.
	PeelSize = 100
)

// aiEndpoint listens for new connections, and handles the game
type aiEndpoint struct {
	bunch    *engine.Bunch
	peelChan chan *peelInfo

	// Fields below are protected by mu
	mu          sync.Mutex
	gameStarted bool
	totalPeels  int
	// Map from username to the Player's AI
	connected map[string]*game
}

func (e *aiEndpoint) connectedPlayers() []string {
	e.mu.Lock()
	defer e.mu.Unlock()
	p := make([]string, len(e.connected))
	i := 0
	for player := range e.connected {
		p[i] = player
		i++
	}
	sort.Strings(p)
	return p
}

// Instead of recording a peel every time it happens, collect all the peels in
// a given interval to reduce frustating race conditions
func (e *aiEndpoint) processPeels() {
	for {
		// Wait for requests to come in
		// TODO(bsprague): *Maybe* dynamically set the sleep time based on client ping times
		time.Sleep(time.Millisecond * 20)
		peelers := make(map[string]potassium.Board)
		for len(e.peelChan) > 0 {
			p := <-e.peelChan
			peelers[p.player] = p.board
		}
		err := e.addSuccessfulPeels(peelers)
		if err != nil {
			log.Printf("Error doing peel stuff: %v", err)
		}
	}
}

func (e *aiEndpoint) startGame() error {
	log.Println("Time to get this thang going")
	e.mu.Lock()
	defer e.mu.Unlock()
	e.peelChan = make(chan *peelInfo, len(e.connected))
	e.gameStarted = true

	// Check for too many players
	if len(e.connected) > MaxPlayers {
		return fmt.Errorf("Hold the fuck up, we only allow %d players, but we have %d players", MaxPlayers, len(e.connected))
	}

	// Hand out the initial tiles.
	tc := engine.StartingTileCount(len(e.connected))

	// tc is on the order of engine.TileScalingFactor, so this is actually kind of large
	for _, game := range e.connected {
		game.mu.Lock()
		game.tiles.Add(e.bunch.TileN(tc))
		game.mu.Unlock()
	}

	go e.processPeels()

	// Now that we've doled out the tiles, send each player a SPLIT request and
	// give them the tiles we just decided on
	for _, game := range e.connected {
		go game.player.Split(context.Background(), func(req potassium.SplitRequest) error {
			game.mu.RLock()
			tiles, err := tilesToWire(game.tiles)
			game.mu.RUnlock()
			if err != nil {
				return err
			}
			if err := req.SetTiles(tiles); err != nil {
				return err
			}

			tl, err := req.NewPlayers(int32(len(e.connected)))
			if err != nil {
				return err
			}

			i := 0
			for name := range e.connected {
				err = tl.Set(i, name)
				if err != nil {
					return fmt.Errorf("sending split: error adding player %s: %v", name, err)
				}
				i++
			}

			return nil
		})
	}

	return nil
}

// addSuccessfulPeels sends PeelSize tiles to all players after one or more
// successful peels are received in a given receive window
func (e *aiEndpoint) addSuccessfulPeels(peelers map[string]potassium.Board) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Increment the total peels and the score of each peeler
	e.totalPeels++
	for peeler := range peelers {
		e.connected[peeler].score += e.totalPeels
	}

	// Because we don't have enough tiles to give everyone, this was the final peel!
	if e.bunch.Count() < len(e.connected)*PeelSize {
		scores := make(map[string]int)
		for _, game := range e.connected {
			scores[game.player.name] = game.score
			// Tell everyone we're done here
			go game.player.GameOver(context.Background(), func(req potassium.Player_gameOver_Params) error { return nil })
		}
		e.gameStarted = false
		if err := db.finishGame(scores); err != nil {
			return fmt.Errorf("failed to save player scores %v to db: %v", scores, err)
		}
		return nil
	}

	// If we're here, it wasn't the final peel
	newTiles := make(map[string]engine.FreqList)
	wireTiles := make(map[string]capnp.TextList)

	// Dole out new tiles from the bunch into each player's hand
	for name, game := range e.connected {
		// Give each player PeelSize new tiles from the bunch
		newTiles[name] = e.bunch.TileN(PeelSize)
		game.mu.Lock()
		game.tiles.Add(newTiles[name])
		game.mu.Unlock()

		tl, err := tilesToWire(newTiles[name])
		if err != nil {
			return err
		}

		wireTiles[name] = tl
	}

	pl, err := peelerTL(peelers)
	if err != nil {
		return err
	}

	// Send out the new tile requests in a separate step, to minimize the delay
	// between players getting the new tiles due to wire conversion
	for name, game := range e.connected {
		go game.player.NewTile(context.Background(), func(req potassium.NewTileRequest) error {
			err = req.SetPeelers(pl)
			if err != nil {
				return fmt.Errorf("sending newTile: error setting peelers %v: %v", pl, err)
			}

			// Send them their new tiles
			err = req.SetLetters(wireTiles[name])
			if err != nil {
				return fmt.Errorf("sending newTile: error setting letters: %v", err)
			}

			return nil
		})
	}

	// Save the peels to the DB
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return fmt.Errorf("saving peel: error making message: %v", err)
	}

	// Create the peel
	p, err := potassium.NewRootPeel(seg)
	if err != nil {
		return fmt.Errorf("saving peel: error making root peel: %v", err)
	}

	vbl, err := p.NewValidBoards(int32(len(peelers)))
	if err != nil {
		return fmt.Errorf("saving peel: error making board list: %v", err)
	}

	i := 0
	for player, board := range peelers {
		pb := vbl.At(i)
		err = pb.SetPlayer(player)
		if err != nil {
			return fmt.Errorf("saving peel: error adding player: %v", err)
		}
		err = pb.SetBoard(board)
		if err != nil {
			return fmt.Errorf("saving peel: error adding board: %v", err)
		}
		i++
	}

	tl, err := p.NewNewTiles(int32(len(newTiles)))
	if err != nil {
		return fmt.Errorf("saving peel: error making new tile list: %v", err)
	}

	i = 0
	for player, letters := range wireTiles {
		e := tl.At(i)
		err = e.SetPlayer(player)
		if err != nil {
			return fmt.Errorf("saving peel: error adding player tiles: %v", err)
		}
		err = e.SetLetters(letters)
		if err != nil {
			return fmt.Errorf("saving peel: error adding tiles: %v", err)
		}
		i++
	}

	return db.addPeel(p)
}

func startAIEndpoint(addr string) (*aiEndpoint, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	e := &aiEndpoint{connected: make(map[string]*game), bunch: engine.NewBunch()}

	// Listen for incoming connections
	go e.listen(l)
	return e, nil
}

func (e *aiEndpoint) listen(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("aiEndpoint: accept: %v\n", err)
		}
		// Handle this connection
		go e.handleConn(c)
	}
}

func (e *aiEndpoint) handleConn(c net.Conn) {
	// Give the connector a handle to the endpoint, so it can add an AI in a sec
	aic := &aiConnector{e: e}
	rc := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(potassium.Server_ServerToClient(aic).Client))
	rc.Wait()
	// End the connection
	aic.drop()
}

type aiConnector struct {
	e    *aiEndpoint
	game *game
}

func (c *aiConnector) Connect(call potassium.Server_connect) error {
	req := call.Params
	resp := call.Results

	name, err := req.Name()
	if err != nil {
		return err
	}

	// If a player is already connected, don't let them connect now
	if c.game != nil {
		log.Printf("%s tried to connect again with name %s\n", c.game.player.name, name)
		resp.SetStatus(potassium.ConnectResponse_Status_yaDoneGoofed)
		return nil
	}

	// TODO(bsprague): Probably limit names to alphanumerics

	if len(name) > 20 {
		// Despite my warnings, some asshat wanted to pick a long name. Well guess
		// what, you don't get to play asshole.
		log.Printf("%s is too long of a name\n", name)
		resp.SetStatus(potassium.ConnectResponse_Status_nameTooLong)
		// TODO(bsprague): IP ban their ass for not following basic instructions,
		// or ban every prefix of this name. Probably not, but damn.
		return nil
	} else if len(name) == 0 {
		// The player sent a blank name. Jesus Christ people, do I have to spell it
		// out for you?
		log.Println("Someone connected with a blank name")
		resp.SetStatus(potassium.ConnectResponse_Status_yaDoneGoofed)
		return nil
	}

	// Lock our connected list
	c.e.mu.Lock()
	defer c.e.mu.Unlock()

	if c.e.gameStarted {
		log.Printf("Player %s tried to connect after the game started", name)
		resp.SetStatus(potassium.ConnectResponse_Status_yaDoneGoofed)
		return nil
	}

	if _, ok := c.e.connected[name]; ok {
		log.Printf("Name %s is already connected\n", name)
		resp.SetStatus(potassium.ConnectResponse_Status_nameAlreadyTaken)
		return nil
	}

	if len(c.e.connected) >= MaxPlayers {
		log.Printf("%s can't join because the game is full\n", name)
		resp.SetStatus(potassium.ConnectResponse_Status_gameIsFull)
		return nil
	}

	p := req.Player()

	// If we're here, we can totally add them
	g := &game{
		e: c.e,
		player: &player{
			Player: p,
			name:   name,
		},
	}
	c.e.connected[name] = g
	c.game = g

	log.Printf("Successfully connected player %s\n", name)
	// Send the player list to web clients
	go sendPlayers()
	resp.SetStatus(potassium.ConnectResponse_Status_success)
	return resp.SetGame(potassium.Game_ServerToClient(g))
}

func (c *aiConnector) drop() {
	// Since we only allow one player to connect per connection, we know exactly
	// who to remove
	c.e.mu.Lock()
	// Only delete them if they successfully connected
	if c.game != nil && c.game.player != nil {
		if _, ok := c.e.connected[c.game.player.name]; ok {
			delete(c.e.connected, c.game.player.name)
			log.Printf("Disconnected player %s\n", c.game.player.name)
			go sendPlayers()
		}
	}
	c.e.mu.Unlock()
}

// This only happens for the initial request
func tilesToWire(tiles engine.FreqList) (capnp.TextList, error) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return capnp.TextList{}, fmt.Errorf("tileToWire: error creating message: %v", err)
	}
	tl, err := capnp.NewTextList(seg, int32(tiles.Count()))
	if err != nil {
		return capnp.TextList{}, fmt.Errorf("tileToWire: error creating capnp.TextList: %v", err)
	}

	i := 0
	for _, tile := range tiles.AsList() {
		err = tl.Set(i, tile)
		if err != nil {
			return capnp.TextList{}, fmt.Errorf("tileToWire: error adding tile: %v", err)
		}
		i++
	}

	return tl, nil
}

func sendPlayers() {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, "player_list.html", globalAIEndpoint.connectedPlayers())
	if err != nil {
		log.Printf("error rendering player list: %v", err)
	}

	err = hub.broadcastJSON(playersAction{
		baseAction: baseAction{Action: "players"},
		HTML:       buf.String(),
	})
	if err != nil {
		log.Printf("error sending player list: %v", err)
	}
}

func peelerTL(peelers map[string]potassium.Board) (capnp.TextList, error) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return capnp.TextList{}, err
	}
	pl, err := capnp.NewTextList(seg, int32(len(peelers)))
	if err != nil {
		return capnp.TextList{}, err
	}

	// Let each player know who peeled
	i := 0
	for peeler := range peelers {
		err = pl.Set(i, peeler)
		if err != nil {
			return capnp.TextList{}, fmt.Errorf("building newTile: error setting peeler %s: %v", peeler, err)
		}
		i++
	}
	return pl, nil
}
