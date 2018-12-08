package game

import (
	"context"
	"net/http"
)

type SpinRequest struct {
	User
}

type SpinResponse struct {
	Total int64  `json:"total"`
	Spins []Spin `json:"spins"`
	User
}

type Spin struct {
	Type  string `json:"type"`
	Total string `json:"total"`
	Stops []int  `json:"stops"`
}

type User struct {
	UID   string
	Chips int64
	Bet   int
}

type service interface {
	Play(ctx context.Context) (Result, error)
}

type GameServer struct {
	service
}

func (gs GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//w.Write()
}

func NewServer() http.Handler {
	return GameServer{}
}
