package machine

type Symbol string

type Reel struct {
	symbols []Symbol
	stopper Stopper
	choices int
}

type Stop struct {
	chosen   []Symbol
	position Position
}

func (r Reel) Spin() Stop {
	pos := r.stopper.Stop()
	return Stop{
		chosen:   r.getSymbols(pos, r.choices),
		position: pos,
	}
}

func (r Reel) getSymbols(position Position, choices int) []Symbol {
	begin := int(position)
	end := begin + choices
	minend := min(end, len(r.symbols))
	return append(r.symbols[begin:minend], r.symbols[:end-minend]...)
}

type Reels []Reel

func (rs Reels) Spin() ([]Stop, error) {
	var stops []Stop
	for _, r := range rs {
		stops = append(stops, r.Spin())
	}
	return stops, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func NewReel(stopper Stopper, symbols []Symbol, choiceLimit int) Reel {
	return Reel{symbols, stopper, choiceLimit}
}
