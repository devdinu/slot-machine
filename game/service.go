package game

import (
	"context"

	"github.com/devdinu/slot_machine/machine"
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
	Compute(ctx context.Context, board Board) (score.Score, error)
}

type Result struct {
	Won   int64
	Stops []int
}

func (s Service) Play(ctx context.Context) (Result, error) {
	spin, err := s.Machine.Spin()
	if err != nil {
		return Result{}, err
	}
	stops := machine.Stops(spin)
	board := stops.GetBoard()
	gameScore, err := s.Compute(ctx, board)
	if err != nil {
		return Result{}, err
	}
	return Result{Stops: stops.GetStopPositions(), Won: gameScore.Points()}, nil
}

func NewService(machine Machine, scorer Scorer) Service {
	return Service{Machine: machine, Scorer: scorer}
}
