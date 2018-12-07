package machine

type Symbol string

type Reel struct {
	symbols []Symbol
	stopper Stopper
}

type Stop struct {
	chosen   []Symbol
	position Position
}

func (r Reel) Spin() Stop {
	pos := r.stopper.Stop()
	return Stop{
		chosen: []Symbol{r.symbols[pos]},
	}
}

type Reels []Reel

func (rs Reels) Spin() ([]Stop, error) {
	var stops []Stop
	for _, r := range rs {
		stops = append(stops, r.Spin())
	}
	return stops, nil
}

func NewReel(stopper Stopper, symbols []Symbol) Reel {
	return Reel{symbols, stopper}
}
