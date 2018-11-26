package main

import (
	"flag"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/pb"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	apiAddr = flag.String("api_addr", ":8081", "RPC server address")
)

func main() {
	var err error

	flag.Parse()

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

	lis, err := net.Listen("tcp", *apiAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterBananaServiceServer(grpcServer, &server{})
	grpcSrv.Serve(lis)

	wrappedGRPC := grpcweb.WrapServer(grpcSrv)
	http.ListenAndServe(*addr, wrappedGRPC)
}
