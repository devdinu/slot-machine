package game

import (
	"context"

	"github.com/devdinu/slot_machine/machine"
	model "github.com/devdinu/slot_machine/models"
	"github.com/devdinu/slot_machine/score"
)

type Board []machine.Symbols

type Service struct {
	Machine
	Scorer
}

type Machine interface {
	Spin() ([]machine.Stop, error)
}

type Scorer interface {
	Compute(ctx context.Context, board model.Board) (score.Score, error)
}

func (s Service) SpinOnce(ctx context.Context, bet int64) (Spin, error) {
	spin, err := s.Machine.Spin()
	if err != nil {
		return Spin{}, err
	}
	stops := machine.Stops(spin)
	board := stops.GetBoard()
	gameScore, err := s.Compute(ctx, board)
	if err != nil {
		return Spin{}, err
	}
	return Spin{Stops: stops.GetStopPositions(), Won: gameScore.Points() * bet, Type: "main"}, nil
}

func (s Service) Play(ctx context.Context, user User) (Result, error) {
	var result Result
	//TODO: implement multiple spins, and pass bet to scorer and multiply
	spin, err := s.SpinOnce(ctx, user.Bet)
	if err != nil {
		return result, nil
	}
	result.Spins = append(result.Spins, spin)
	result.TotalWin += spin.Won
	result.User.Chips -= user.Bet
	return result, nil
}

func NewService(machine Machine, scorer Scorer) Service {
	return Service{Machine: machine, Scorer: scorer}
}
