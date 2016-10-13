// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	txt "text/template"
	"time"

	"github.com/bcspragu/Bananagrama/engine"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
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
	dict             engine.Dictionary
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

	dict = engine.NewDictionary(f)
	rand.Seed(time.Now().UnixNano())

	globalAIEndpoint, err = startAIEndpoint(*apiAddr)
	if err != nil {
		log.Fatal("AI RPC endpoint failed to start:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/pass", servePass).Methods("POST")
	r.HandleFunc("/startGame", serveGame).Methods("GET")
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	if *env == "dev" {
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	}

	http.Handle("/", r)

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveGame(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("key") == "" {
		fmt.Fprint(w, "key not set")
		return
	}
	if err := globalAIEndpoint.startGame(); err != nil {
		fmt.Fprintf(w, "failed to start game: %v", err)
	}
}
