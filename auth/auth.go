package auth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bcspragu/Bananagrama/banana"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const authMDKey = "internal-player-id"

type Client struct {
	pk *ecdsa.PrivateKey
}

func New(pk *ecdsa.PrivateKey) *Client {
	return &Client{pk: pk}
}

func (c *Client) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authentication for the Register RPC.
	if info.FullMethod == "/BananaService/Register" {
		return handler(ctx, req)
	}

	// Otherwise, look for the token.
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		log.Printf("failed to get token from context: %v", err)
		return nil, errors.New("unauthorized")
	}

	pID, err := c.verify(tkn)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return nil, errors.New("unauthorized")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}

	// Add the player ID to the metadata and update the context.
	md.Set(authMDKey, string(pID))
	ctx = metadata.NewIncomingContext(ctx, md)

	return handler(ctx, req)
}

func (c Client) StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Look for the token.
	tkn, err := tokenFromContext(ss.Context())
	if err != nil {
		log.Printf("failed to get token from context: %v", err)
		return errors.New("unauthorized")
	}

	pID, err := c.verify(tkn)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return errors.New("unauthorized")
	}

	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		md = metadata.New(map[string]string{})
	}

	// Add the player ID to the metadata and update the context.
	md.Set(authMDKey, string(pID))

	// Add the player ID to the metadata.
	if err := ss.SendHeader(md); err != nil {
		log.Printf("failed to add player ID to metadata: %v", err)
		return errors.New("unauthorized")
	}

	return handler(srv, ss)
}

func (c *Client) PlayerIDFromContext(ctx context.Context) (banana.PlayerID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata in context")
	}

	v := md.Get(authMDKey)
	if len(v) == 0 {
		return "", errors.New("no player ID in metadata")
	}
	if len(v) > 1 {
		return "", fmt.Errorf("%d player IDs in metadata", len(v))
	}

	return banana.PlayerID(v[0]), nil
}

func tokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata in context")
	}

	cookies := md.Get("cookie")

	if len(cookies) != 1 {
		return "", fmt.Errorf("invalid number of cookie entries %d", len(cookies))
	}

	var tkn string
	for _, cookie := range strings.Split(cookies[0], "; ") {
		if strings.HasPrefix(cookie, "player-token=") {
			tkn = strings.TrimPrefix(cookie, "player-token=")
			break
		}
	}

	if tkn == "" {
		return "", errors.New("token not found in cookies")
	}

	return tkn, nil
}

func (c *Client) verify(tkn string) (banana.PlayerID, error) {
	token, err := jwt.ParseWithClaims(tkn, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		m, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if m.Name != "ES256" {
			return nil, fmt.Errorf("unexpected signing name %q", m.Name)
		}
		return &c.pk.PublicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return "", errors.New("token was invalid")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("unexpected claims type %T", token.Claims)
	}

	if err := claims.Valid(); err != nil {
		return "", fmt.Errorf("invalid token claims: %v", err)
	}

	if claims.Issuer != "bananagrams" {
		return "", fmt.Errorf("invalid issuer %q", claims.Issuer)
	}

	return banana.PlayerID(claims.Subject), nil
}

func (c *Client) Sign(pID banana.PlayerID) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Subject:  string(pID),
		IssuedAt: time.Now().Unix(),
		Issuer:   "bananagrams",
	})

	tknStr, err := tkn.SignedString(c.pk)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tknStr, nil
}
