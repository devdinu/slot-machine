package machine

type Stopper interface {
	Stop() Position
}

type Machine struct {
	Reels
}

func (m Machine) Spin() ([]Stop, error) {
	return m.Reels.Spin()
}

func NewMachine(stopper Stopper, syms []Symbol, totalReels int, totalChoiceSymbols int) Machine {
	var reels []Reel
	for i := 0; i < totalReels; i++ {
		reels = append(reels, NewReel(stopper, syms, totalChoiceSymbols))
	}
	return Machine{Reels: reels}
}
