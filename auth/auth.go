package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/dgrijalva/jwt-go"
)

type Token string

type Client struct {
	pk *ecdsa.PrivateKey
}

func New() (*Client, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	return &Client{pk: pk}, nil
}

func (c *Client) Verify(tkn Token) (banana.PlayerID, error) {
	token, err := jwt.Parse(string(tkn), func(token *jwt.Token) (interface{}, error) {
		m, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if m.Name != "ES256" {
			return nil, fmt.Errorf("Unexpected signing name %q", m.Name)
		}
		return c.pk, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return "", errors.New("token was invalid")
	}

	return "", errors.New("IMPLEMENT ME")
}

func (c *Client) Sign() {

}
