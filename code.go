package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var code string

func loadCode() error {
	c, err := ioutil.ReadFile("code")
	if err != nil {
		return fmt.Errorf("Error reading password: %s", err)
	}

	code = strings.TrimSpace(string(c))
	return nil
}
