package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bcspragu/Bananagrama/auth"
	"github.com/bcspragu/Bananagrama/banana"
	"github.com/bcspragu/Bananagrama/memdb"
	"github.com/bcspragu/Bananagrama/pb"
	"github.com/bcspragu/Bananagrama/srv"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

var (
	addr        = flag.String("addr", ":8080", "http service address")
	apiAddr     = flag.String("api_addr", ":8079", "RPC server address")
	dictPath    = flag.String("dict", "", "path to the dictionary file to use")
	authKeyPath = flag.String("auth_key_path", "auth.key", "path to the PEM-encoded ECDSA key to use for signing/verifying JWT tokens")
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

	pk, err := loadOrGenerateKey(*authKeyPath)

	authClient := auth.New(pk)

	grpcSrv := grpc.NewServer(
		grpc.UnaryInterceptor(authClient.UnaryInterceptor),
		grpc.StreamInterceptor(authClient.StreamInterceptor),
	)
	server, err := srv.New(r, authClient, memdb.New(r), dict)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}

	pb.RegisterBananaServiceServer(grpcSrv, server)
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("grpc.Serve: %v", err)
		}
	}()

	http.ListenAndServe(*addr, grpcweb.WrapServer(grpcSrv))
}

func loadDict(fn string) (banana.Dictionary, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return banana.NewDictionary(f)
}

func loadOrGenerateKey(keyPath string) (*ecdsa.PrivateKey, error) {
	f, err := os.Open(keyPath)
	if os.IsNotExist(err) {
		return generateAndSaveKey(keyPath)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %v", err)
	}
	defer f.Close()

	// File exists, load from it.
	return loadKey(f)
}

func generateAndSaveKey(keyPath string) (*ecdsa.PrivateKey, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	dat, err := x509.MarshalECPrivateKey(pk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal key: %v", err)
	}

	if err := ioutil.WriteFile(keyPath, dat, 0600); err != nil {
		return nil, fmt.Errorf("failed to write auth key: %v", err)
	}

	return pk, nil
}

func loadKey(r io.Reader) (*ecdsa.PrivateKey, error) {
	dat, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read auth key: %v", err)
	}

	pk, err := x509.ParseECPrivateKey(dat)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key: %v", err)
	}

	return pk, nil
}
