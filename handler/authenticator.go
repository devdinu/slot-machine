package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/devdinu/slot_machine/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func Authenticator(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		_, err := authenticate(w.Header().Get("Authorization"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
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
