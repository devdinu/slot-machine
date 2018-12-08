package game

import (
	"context"
	"encoding/json"
	"net/http"
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

func (gs gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

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

func NewServer() http.Handler {
	return gameServer{}
}
