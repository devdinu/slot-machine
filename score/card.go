package score

import model "github.com/devdinu/slot_machine/models"

type symbolScore []int

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
