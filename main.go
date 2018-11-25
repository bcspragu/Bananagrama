package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	txt "text/template"
	"time"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
	"github.com/gorilla/securecookie"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	apiAddr = flag.String("api_addr", ":8081", "RPC server address")
	env     = flag.String("env", "dev", "Which environment to run in")

	templates        *template.Template
	txtTmpl          *txt.Template
	globalAIEndpoint *aiEndpoint
	hub              *Hub
	db               datastore
	s                *securecookie.SecureCookie
	dict             banana.Dictionary
)

func main() {
	var err error

	flag.Parse()
	hub = newHub()
	go hub.run()
	templates = template.Must(template.New("templates").Funcs(template.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
	}).ParseGlob("templates/*"))
	txtTmpl = txt.Must(txt.New("js").Funcs(txt.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
	}).ParseGlob("js/*"))

	if err := loadPass(); err != nil {
		panic(err)
	}

	if err := loadCode(); err != nil {
		panic(err)
	}

	if s, err = initKeys(); err != nil {
		panic(err)
	}

	db, err = initDB("bananagrama.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	f, err := os.Open("dict.txt")
	if err != nil {
		panic(err)
	}

	dict = banana.NewDictionary(f)
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterBananaServiceServer(grpcServer, &server{})
	grpcSrv.Serve(lis)

	wrappedGRPC := grpcweb.WrapServer(grpcSrv)
	http.ListenAndServe(*addr, wrappedGRPC)
}

type server struct{}

func (s *server) NewGame(ctx context.Context, req *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *server) JoinGame(req *pb.JoinGameRequest, stream pb.BananaService_JoinGameServer) error {
	return nil
}

func (s *server) Dump(ctx context.Context, req *pb.DumpRequest) (*pb.DumpResponse, error) {
	return nil, errors.New("not implemented")
}
