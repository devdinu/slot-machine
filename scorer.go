package machine

type Scorer struct {
	payLines []Line
	//card     Scorecard
}

type Score struct {
}

func (s *Scorer) Compute() (Score, error) {
	return Score{}, nil
}
