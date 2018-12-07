package score

type Symbol string

type symbolScore []int

type scoreCard struct {
	symbolScores map[Symbol]symbolScore
}

func (sc scoreCard) score(sym Symbol, occ int) int64 {
	score, ok := sc.symbolScores[sym]
	if !ok || len(score) < occ {
		return -1
	}
	return int64(score[occ])
}
