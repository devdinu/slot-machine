package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/devdinu/slot_machine/config"
	model "github.com/devdinu/slot_machine/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func Authenticator(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		//TODO: parse header with Bearer
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
	ctx := context.WithValue(parent, model.UserBetKey, claims.Bet)
	ctx = context.WithValue(ctx, model.UserIDKey, claims.UID)
	ctx = context.WithValue(ctx, model.UserChipsKey, claims.Chips)
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

type Refresher struct {
	http.ResponseWriter
}

func (r *Refresher) Refresh(bet, chips int64, uid string) *jwt.Token {
	claims := &UserClaim{}
	claims.Bet = bet
	claims.Chips = chips
	claims.UID = uid
	claims.ExpiresAt = time.Now().Add(config.AuthTokenExpiryMinutes()).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}
