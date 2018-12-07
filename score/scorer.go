package score

import (
	"context"
	"errors"
)

type Line []location
type Symbols []Symbol
type Board []Symbols

type Scorer struct {
	paylines []Line
	card     scoreCard
}

type occurence struct {
	sym   Symbol
	count int
}
type location struct {
	row int
	col int
}

type Score struct {
	won int64
}

func (b Board) get(l location) Symbol {
	return b[l.row][l.col]
}

func (b Board) empty() bool {
	return len(b) == 0 || len(b[0]) == 0
}

func (s *Scorer) Compute(ctx context.Context, board Board) (Score, error) {
	var score Score
	for _, pl := range s.paylines {
		occ, err := s.findOccurrences(pl, board)
		if err != nil {
			return Score{}, err
		}
		score.won += s.card.score(occ.sym, occ.count)
		//fmt.Println("score-----", score.won, occ.sym, occ.count)
	}
	return score, nil
}

func (s *Scorer) findOccurrences(line Line, board Board) (occurence, error) {
	if board.empty() || len(line) == 0 {
		return occurence{}, errors.New("Invalid Board")
	}
	first := board.get(line[0])
	count := 0
	for _, loc := range line {
		if board.get(loc) == first {
			count++
		} else {
			break
		}
	}
	return occurence{first, count}, nil
}
