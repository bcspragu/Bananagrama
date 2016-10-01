package main

import "github.com/gopherjs/gopherjs/js"

//go:generate gopherjs build --output=../js/gopher.js
func main() {
	js.Global.Set("Bananarama", map[string]interface{}{
	// Add some functions here for decoding the encoded map
	})
}

func GetPlayback(base64Playback string) *js.Object {
	/*
		TODO(bsprague): Add some notion of a playback
		var p engine.Playback
		buf := bytes.NewReader([]byte(base64Playback))
		r := base64.NewDecoder(base64.StdEncoding, buf)
		gob.NewDecoder(r).Decode(&p)
		return js.MakeWrapper(&p)
	*/
	return js.MakeWrapper(nil)
}
