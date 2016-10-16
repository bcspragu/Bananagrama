package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type indexData struct {
	JS template.JS
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	var jsBuf bytes.Buffer

	err := txtTmpl.ExecuteTemplate(&jsBuf, "game.js", struct {
		WSHost string
		Dev    bool
	}{
		r.Host,
		*env == "dev",
	})

	err = templates.ExecuteTemplate(w, "index.html", struct {
		JS      template.JS
		Players []string
	}{
		template.JS(jsBuf.String()),
		globalAIEndpoint.connectedPlayers(),
	})
	if err != nil {
		fmt.Println(err)
	}
}
