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

type Config struct {
	ReelsOfSymbols []Symbols
	Rows           int
}

func NewMachine(stopper Stopper, cfg Config) Machine {
	var reels Reels
	for _, reel := range cfg.ReelsOfSymbols {
		reels = append(reels, NewReel(stopper, reel, cfg.Rows))
	}
	return Machine{Reels: reels}
}
