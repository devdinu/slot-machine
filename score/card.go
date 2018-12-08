package score

import model "github.com/devdinu/slot_machine/models"

type symbolScore []int

// holds score for symbol and occurence
// symbol: [0, 10, 20, 30], occurence for 1,2,3,4 in order
type scoreCard struct {
	symbolScores map[model.Symbol]symbolScore
}

func (sc scoreCard) score(sym model.Symbol, occ int) int64 {
	score, ok := sc.symbolScores[sym]
	if !ok || len(score) < occ {
		return 0
	}
	return int64(score[occ-1])
}
