package score

import (
	"context"
	"errors"
	"fmt"

	model "github.com/devdinu/slot_machine/models"
)

type Scorer struct {
	paylines []model.Line
	card     scoreCard
	scatter  model.Symbol
}

type occurence struct {
	sym   model.Symbol
	count int
}

type Score struct {
	Won int64
}

func (s *Score) Points() int64 {
	return s.Won
}

func (s Scorer) Compute(ctx context.Context, board model.Board) (Score, error) {
	var score Score
	for _, pl := range s.paylines {
		occ, err := s.findOccurrences(pl, board)
		if err != nil {
			fmt.Println("received an error: ", err)
			return Score{}, err
		}
		score.Won += s.card.score(occ.sym, occ.count)
		fmt.Println("score by occ -----", score.Won, occ.sym, occ.count, &score)
	}
	score.Won += s.scatterScore(ctx, board)
	return score, nil
}

func (s Scorer) scatterScore(ctx context.Context, board model.Board) int64 {
	scatterCount := 0
	for _, row := range board {
		for _, sym := range row {
			if sym == s.scatter {
				scatterCount += 1
			}
		}
	}
	return s.card.score(s.scatter, scatterCount)
}

func (s Scorer) findOccurrences(line model.Line, board model.Board) (occurence, error) {
	if board.Empty() || len(line) == 0 {
		return occurence{}, errors.New("Invalid Board")
	}
	first := board.Get(line[0])
	count := 0
	for _, loc := range line {
		if board.Get(loc) == first {
			count++
		} else {
			break
		}
	}
	return occurence{first, count}, nil
}

func NewScorer(scatter model.Symbol) Scorer {
	return Scorer{
		scatter: scatter,
	}
}
