package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var gameJS string

func servePass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Valid bool
		Game  string
		Error string
		JS    string
	}{}

	if r.PostFormValue("password") == password {
		var htmlBuf, jsBuf bytes.Buffer

		err := txtTmpl.ExecuteTemplate(&jsBuf, "game.js", struct {
			WSHost string
			Dev    bool
		}{
			r.Host,
			*env == "dev",
		})
		if err != nil {
			fmt.Printf("Error reading JS: %s\n", err)
		}

		err = templates.ExecuteTemplate(&htmlBuf, "game.html", nil)
		if err != nil {
			fmt.Println(err)
		}

		data.Valid = true
		data.Game = htmlBuf.String()
		data.JS = jsBuf.String()
		if err := setCookie(w, password); err != nil {
			fmt.Println(err)
		}
	} else {
		data.Error = "Invalid password"
	}
	respString, _ := json.Marshal(data)
	fmt.Fprint(w, string(respString))
}

func setCookie(w http.ResponseWriter, pass string) error {
	encoded, err := s.Encode("pass", pass)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "pass",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	return err
}
