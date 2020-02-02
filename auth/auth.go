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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type playerIDKey struct{}

func ContextWithPlayerID(parent context.Context, pID banana.PlayerID) context.Context {
	return context.WithValue(parent, playerIDKey{}, pID)
}

func PlayerIDFromContext(ctx context.Context) (banana.PlayerID, bool) {
	pID, ok := ctx.Value(playerIDKey{}).(banana.PlayerID)
	if !ok {
		return "", false
	}
	return pID, true
}

type DB interface {
	Player(id banana.PlayerID) (*banana.Player, error)
}

type Client struct {
	pk *ecdsa.PrivateKey
	db DB
}

func New(pk *ecdsa.PrivateKey, db DB) *Client {
	return &Client{pk: pk, db: db}
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
		return nil, status.Error(codes.Unauthenticated, "couldn't get token from context")
	}

	pID, err := c.verify(tkn)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token in context")
	}

	if _, err := c.db.Player(pID); err != nil {
		log.Printf("failed to load player from ID in context: %v", err)
		return nil, status.Error(codes.Unauthenticated, "player not found")
	}

	return handler(ContextWithPlayerID(ctx, pID), req)
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func (c Client) StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Look for the token.
	tkn, err := tokenFromContext(ss.Context())
	if err != nil {
		log.Printf("failed to get token from context: %v", err)
		return status.Error(codes.Unauthenticated, "no token in context")
	}

	pID, err := c.verify(tkn)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return status.Error(codes.Unauthenticated, "invalid token in context")
	}

	if _, err := c.db.Player(pID); err != nil {
		log.Printf("failed to load player from ID in context: %v", err)
		return status.Error(codes.Unauthenticated, "player not found")
	}

	newSS := &serverStream{
		ServerStream: ss,
		ctx:          ContextWithPlayerID(ss.Context(), pID),
	}

	return handler(srv, newSS)
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
