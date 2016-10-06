package main

import (
	"context"
	"log"
	"net"
	"sync"

	"zombiezen.com/go/capnproto2/rpc"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/bcspragu/Bananagrama/potassium"
)

const (
	MaxPlayers = 8
)

// aiEndpoint listens for new connections, and handles the game
type aiEndpoint struct {
	bunch *engine.Bunch

	// Fields below are protected by mu
	mu sync.Mutex
	// Map from username to the Player's AI
	connected map[string]*game
}

// sendNewTiles sends a tile to all players after a successful peel
func (e *aiEndpoint) sendNewTiles(peeler string) {
	for _, game := range e.connected {
		// TODO(bsprague): Make sure we have enough letters first in the bunch,
		// otherwise end the game. And log and keep track and stuff
		go game.player.NewTile(context.Background(), func(req potassium.NewTileRequest) error {
			tile := e.bunch.Tile()
			req.SetLetter(tile.String())
			game.tiles.Inc(tile)
			req.SetPeeler(peeler)
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
	e      *aiEndpoint
	player string
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
	if c.player != "" {
		log.Printf("%s tried to connect again with name %s\n", c.player, name)
		resp.SetStatus(potassium.ConnectResponse_Status_yaDoneGoofed)
		return nil
	}

	// TODO(bsprague): Probably limit names to alphanumerics. Oh, and log when
	// someone forces a non-success response

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

	player := req.Player()

	// If we're here, we can totally add them
	g := &game{
		e:      c.e,
		player: player,
		name:   name,
	}
	c.e.connected[name] = g
	c.player = name
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
	delete(c.e.connected, c.player)
	c.e.mu.Unlock()
	log.Printf("Disconnected player %s\n", c.player)
}
