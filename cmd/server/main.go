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
	"github.com/bcspragu/Bananagrama/memdb"
	"github.com/bcspragu/Bananagrama/pb"
	"github.com/bcspragu/Bananagrama/srv"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

var (
	addr     = flag.String("addr", ":8080", "http service address")
	apiAddr  = flag.String("api_addr", ":8079", "RPC server address")
	dictPath = flag.String("dict", "", "path to the dictionary file to use")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *apiAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dict, err := loadDict(*dictPath)
	if err != nil {
		log.Fatalf("failed to load dictionary %q: %v", *dictPath, err)
	}

	grpcSrv := grpc.NewServer()
	server := srv.New(r, memdb.New(r, dict), dict)
	pb.RegisterBananaServiceServer(grpcSrv, server)
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("grpc.Serve: %v", err)
		}
	}()

	wrappedGRPC := grpcweb.WrapServer(grpcSrv)

	http.ListenAndServe(*addr, wrappedGRPC)
}

func loadDict(fn string) (banana.Dictionary, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return banana.NewDictionary(f)
}
