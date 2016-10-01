package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	storedPass, err := readCookie(r)
	if err == nil && storedPass == password {
		err = renderGame(w, r)
		if err != nil {
			fmt.Printf("error rendering game: %v\n", err)
		}
		return
	} else if err != nil {
		fmt.Printf("error reading cookie: %v\n", err)
	}
	err = templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func readCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("pass")
	if err == nil {
		var val string
		err = s.Decode("pass", cookie.Value, &val)
		if err == nil {
			return val, nil
		}
		return "", err
	}
	return "", err
}

func renderGame(w http.ResponseWriter, r *http.Request) error {
	var jsBuf bytes.Buffer

	err := txtTmpl.ExecuteTemplate(&jsBuf, "game.js", struct {
		WSHost string
		Dev    bool
	}{
		r.Host,
		*env == "dev",
	})
	if err != nil {
		return err
	}
	err = templates.ExecuteTemplate(w, "game.html", struct {
		WSHost   string
		FullLoad bool
		JS       template.JS
	}{
		r.Host,
		true,
		template.JS(jsBuf.String()),
	})
	if err != nil {
		return err
	}
	return nil
}
