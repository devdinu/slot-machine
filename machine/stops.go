package machine

type Symbol string
type Stops []Stop
type Symbols []Symbol

type Stop struct {
	Symbols
	Position
}

func (stp Stops) GetBoard() []Symbols {
	board := make([]Symbols, len(stp))
	for i, s := range stp {
		board[i] = s.Symbols
	}
	return board
}

func (stp Stops) GetStopPositions() []int {
	positions := make([]int, len(stp))
	for i, s := range stp {
		positions[i] = int(s.Position)
	}
	return positions
}
