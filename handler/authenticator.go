package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/devdinu/slot_machine/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type CtxKey string

var UserBetKey = CtxKey("bet")
var UserIDKey = CtxKey("uid")
var UserChipsKey = CtxKey("chips")

func Authenticator(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		claims, err := authenticate(authToken)
		if err != nil {
			fmt.Println("Unauthorized request", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := addContextInfo(r.Context(), claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(f)
}

func addContextInfo(parent context.Context, claims *UserClaim) context.Context {
	ctx := context.WithValue(parent, UserBetKey, claims.Bet)
	ctx = context.WithValue(ctx, UserIDKey, claims.UID)
	ctx = context.WithValue(ctx, UserChipsKey, claims.Chips)
	return ctx
}

type UserClaim struct {
	Bet   int64  `json:"bet"`
	Chips int64  `json:"chips"`
	UID   string `json:"uid"`
	jwt.StandardClaims
}

func tokenParser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Invalid signing method: %s, expected HMAC", token.Method)
	}
	return config.AuthSecret(), nil
}

func authenticate(authToken string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(authToken, &UserClaim{}, tokenParser)
	if err != nil {
		return nil, fmt.Errorf("Token Parse Failure: %v", err)
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok || !token.Valid {
		return nil, errors.New("Token Invalid")
	}
	return claims, nil
}
