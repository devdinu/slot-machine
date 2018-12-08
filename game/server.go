package game

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	model "github.com/devdinu/slot_machine/models"
	jwt "github.com/dgrijalva/jwt-go"
)

type SpinRequest struct {
	User
}

type Result struct {
	Spins    `json:"spins"`
	TotalWin int64 `json:"total"`
	User     `json:"user"`
	Token    string `json:"jwt_token,omitempty"`
}

type Spin struct {
	Type  string `json:"type"`
	Won   int64  `json:"total"`
	Stops []int  `json:"stops"`
}

type User struct {
	UID   string `json:"uid"`
	Chips int64  `json:"chips"`
	Bet   int64  `json:"bet"`
}
type Spins []Spin

type service interface {
	Play(context.Context, User) (Result, error)
}
type tokenRefresher interface {
	Refresh(bet, chips int64, uid string) *jwt.Token
}

type gameServer struct {
	service
	tokenRefresher
}

func getUserInfo(ctx context.Context) (User, error) {
	bet, ok1 := ctx.Value(model.UserBetKey).(int64)
	chips, ok2 := ctx.Value(model.UserChipsKey).(int64)
	uid, ok3 := ctx.Value(model.UserIDKey).(string)
	if !ok1 || !ok2 || !ok3 {
		return User{}, errors.New("Invalid Request")
	}
	return User{Bet: bet, Chips: chips, UID: uid}, nil
}

func (gs gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := getUserInfo(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := gs.service.Play(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := gs.tokenRefresher.Refresh(user.Bet, res.User.Chips, res.User.UID)
	res.Token, _ = token.SigningString()
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewServer(svc service, refresher tokenRefresher) http.Handler {
	return gameServer{svc, refresher}
}
