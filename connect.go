package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	capnp "zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/rpc"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
)

const (
	// How many people can join in a single game
	MaxPlayers = 8
	// How many tiles a player receives on a dump
	DumpSize = 3
)

// aiEndpoint listens for new connections, and handles the game
type aiEndpoint struct {
	bunch *engine.Bunch

	// Fields below are protected by mu
	mu         sync.Mutex
	totalPeels int
	// Map from username to the Player's AI
	connected map[string]*game
}

func (e *aiEndpoint) startGame() error {
	log.Println("Time to get this thang going")
	e.mu.Lock()
	defer e.mu.Unlock()

	// Check for too many players
	if len(e.connected) > MaxPlayers {
		return fmt.Errorf("Hold the fuck up, we only allow %d players, but we have %d players", MaxPlayers, len(e.connected))
	}

	// Hand out the initial tiles. Go round robin, even if it's less efficient and unnecessary
	tc := engine.StartingTileCount(len(e.connected))

	// tc is on the order of engine.TileScalingFactor, so this is actually kind of large
	for i := 0; i < tc; i++ {
		for _, game := range e.connected {
			game.AddTile(e.bunch.Tile())
		}
	}

	// Now that we've doled out the tiles, send each player a SPLIT request and
	// give them the tiles we just decided on
	for _, game := range e.connected {
		go game.player.Split(context.Background(), func(req potassium.SplitRequest) error {
			tiles, err := tilesToWire(game.tiles)
			if err != nil {
				return err
			}
			if err := req.SetTiles(tiles); err != nil {
				return err
			}

			e.mu.Lock()
			tl, err := req.NewPlayers(int32(len(e.connected)))
			if err != nil {
				return err
			}

			i := 0
			for name := range e.connected {
				tl.Set(i, name)
				i++
			}
			e.mu.Unlock()

			return nil
		})
	}

	return nil
}

// addSuccessfulPeel sends a tile to all players after a successful peel
func (e *aiEndpoint) addSuccessfulPeel(peeler string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Increment the total peels and their score
	e.totalPeels++
	e.connected[peeler].score += e.totalPeels

	// This was the final peel!
	if e.bunch.Count() < len(e.connected) {
		// TODO(bsprague) Add other endgame stuff (ie datastore things)
		for _, game := range e.connected {
			// Tell everyone we're done here
			go game.player.GameOver(context.Background(), func(req potassium.Player_gameOver_Params) error { return nil })
		}
		return
	}
	for _, game := range e.connected {
		// TODO(bsprague): Log stuff, add
		// Give each player a new tile
		go game.player.NewTile(context.Background(), func(req potassium.NewTileRequest) error {
			req.SetPeeler(peeler)

			tile := e.bunch.Tile()
			req.SetLetter(tile.String())
			game.tiles.Inc(tile) // Add the tile to their hand
			return nil
		})
	}
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
	// The player connected via this connection
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
	c.e.mu.Unlock()

	log.Printf("Successfully connected player %s\n", name)
	resp.SetStatus(potassium.ConnectResponse_Status_success)
	resp.SetGame(potassium.Game_ServerToClient(g))

	return nil
}

func (c *aiConnector) drop() {
	// Since we only allow one player to connect per connection, we know exactly
	// who to remove
	c.e.mu.Lock()
	delete(c.e.connected, c.game.player.name)
	c.e.mu.Unlock()
	log.Printf("Disconnected player %s\n", c.game.player.name)
}

// This only happens for the initial request
func tilesToWire(tiles engine.FreqList) (capnp.TextList, error) {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return capnp.TextList{}, err
	}
	tl, err := capnp.NewTextList(seg, int32(tiles.Count()))
	if err != nil {
		return capnp.TextList{}, err
	}

	i := 0
	for _, tile := range tiles.AsList() {
		tl.Set(i, tile)
		i++
	}

	return tl, nil
}
