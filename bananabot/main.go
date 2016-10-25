package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/bcspragu/Bananagrama/potassium"
	"zombiezen.com/go/capnproto2/rpc"
)

type player struct {
	game potassium.Game
}

var (
	addr = flag.String("addr", "localhost:8081", "http service address")
	name = flag.String("name", "bananabot", "player name")
)

func main() {
	flag.Parse()
	p := &player{}
	psc := potassium.Player_ServerToClient(p)

	c, err := net.Dial("tcp", *addr)
	if err != nil {
		panic(err)
	}
	conn := rpc.NewConn(rpc.StreamTransport(c))
	s := potassium.Server{Client: conn.Bootstrap(context.TODO())}
	resp, err := s.Connect(context.TODO(), func(r potassium.ConnectRequest) error {
		if err := r.SetName(*name); err != nil {
			return err
		}
		if err := r.SetPlayer(psc); err != nil {
			return err
		}
		return nil
	}).Struct()
	if err != nil {
		log.Printf("Failed to connect: %v", err)
	}

	p.game = resp.Game()
	log.Println("Connected, waiting for game to start")
	conn.Wait()
}

func (p *player) Split(call potassium.Player_split) error {
	log.Println("Game is starting!")
	// Ignoring errors for clarity, definitely don't do this in your real code
	tiles, _ := call.Params.Tiles()
	players, _ := call.Params.Players()

	log.Printf("Received %d tiles\n", tiles.Len())

	for i := 0; i < players.Len(); i++ {
		p, _ := players.At(i)
		log.Printf("Playing against %s", p)
	}

	// This would be a good time to start computing your board and stuff
	// p.game.Peel(...) and stuff
	return nil
}

func (p *player) NewTile(call potassium.Player_newTile) error {
	// Add the new tiles to our board and peel again
	return nil
}

func (p *player) DumpNotice(call potassium.Player_dumpNotice) error {
	return nil
}

func (p *player) GameOver(call potassium.Player_gameOver) error {
	return nil
}
