package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var code, password string

func loadCode() error {
	c, err := ioutil.ReadFile("code")
	if err != nil {
		return fmt.Errorf("Error reading password: %s", err)
	}

	code = strings.TrimSpace(string(c))
	return nil
}

func loadPass() error {
	pass, err := ioutil.ReadFile("password")
	if err != nil {
		return fmt.Errorf("Error reading password: %s", err)
	}

	password = strings.TrimSpace(string(pass))
	return nil
}
