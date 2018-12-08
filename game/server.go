package game

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devdinu/slot_machine/handler"
)

type SpinRequest struct {
	User
}

type Result struct {
	Spins    `json:"spins"`
	TotalWin int64 `json:"total"`
	User     `json:"user"`
}

type Spin struct {
	Type  string `json:"type"`
	Won   int64  `json:"total"`
	Stops []int  `json:"stops"`
}

type User struct {
	UID   string
	Chips int64
	Bet   int64
}
type Spins []Spin

type service interface {
	Play(context.Context, User) (Result, error)
}

type gameServer struct {
	service
}

func getUserInfo(ctx context.Context) User {
	bet, ok1 := ctx.Value(handler.UserBetKey).(int64)
	chips, ok2 := ctx.Value(handler.UserChipsKey).(int64)
	uid, ok3 := ctx.Value(handler.UserIDKey).(string)
	fmt.Println(ok1, ok2, ok3)
	if !ok1 || !ok2 || !ok3 {
		return User{}
	}
	return User{Bet: bet, Chips: chips, UID: uid}
}

func (gs gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := getUserInfo(r.Context())
	fmt.Println("proccessing request for", user)

	res, err := gs.service.Play(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewServer(svc service) http.Handler {
	return gameServer{svc}
}
