package score

import (
	"context"
	"errors"
	"sync"

	model "github.com/devdinu/slot_machine/models"
)

type Scorer struct {
	paylines []model.Line
	card     ScoreCard
	scatter  model.Symbol
	wild     model.Symbol
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
	var wg sync.WaitGroup
	wg.Add(len(s.paylines) + 1) //+1 for scatter score
	scoreChan := make(chan int64, len(s.paylines)+1)
	errChan := make(chan error, len(s.paylines))

	for _, pl := range s.paylines {
		go s.lineScore(ctx, pl, board, &wg, scoreChan, errChan)
	}

	//could use Socre with locks to avoid channels and have simplicity
	go s.scatterScore(ctx, board, &wg, scoreChan, errChan)
	wg.Wait()
	close(scoreChan)
	close(errChan)
	if err, ok := <-errChan; ok && err != nil {
		return Score{}, err
	}
	for cscore := range scoreChan {
		score.Won += cscore
	}
	return score, nil
}

func (s Scorer) lineScore(ctx context.Context, line model.Line, board model.Board, wg *sync.WaitGroup, scoreChan chan<- int64, errChan chan error) {
	defer wg.Done()

	occ, err := s.findOccurrences(line, board)
	if err != nil {
		errChan <- err
	}
	scoreChan <- s.card.score(occ.sym, occ.count)
}

func (s Scorer) scatterScore(ctx context.Context, board model.Board, wg *sync.WaitGroup, scoreChan chan<- int64, errChan chan error) {
	defer wg.Done()

	scatterCount := 0
	for _, row := range board {
		for _, sym := range row {
			if sym == s.scatter {
				scatterCount += 1
			}
		}
	}
	scoreChan <- s.card.score(s.scatter, scatterCount)
}

func (s Scorer) findOccurrences(line model.Line, board model.Board) (occurence, error) {
	if board.Empty() || len(line) == 0 {
		return occurence{}, errors.New("Invalid Board")
	}
	first := board.Get(line[0])
	count := 0
	for _, loc := range line {
		currSym := board.Get(loc)
		if currSym == first || currSym == s.wild {
			count++
		} else {
			break
		}
	}
	return occurence{first, count}, nil
}

type Config struct {
	Wild     model.Symbol
	Scatter  model.Symbol
	Paylines []model.Line
	ScoreCard
}

func NewScorer(cfg Config) Scorer {
	return Scorer{
		scatter:  cfg.Scatter,
		wild:     cfg.Wild,
		paylines: cfg.Paylines,
		card:     cfg.ScoreCard,
	}
}
