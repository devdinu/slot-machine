package machine

import model "github.com/devdinu/slot_machine/models"

type Symbol string
type Symbols []Symbol

type Stop struct {
	Symbols
	Position
}

type Stops []Stop

func (stp Stops) GetStopPositions() []int {
	positions := make([]int, len(stp))
	for i, s := range stp {
		positions[i] = int(s.Position)
	}
	return positions
}

func (stp Stops) GetBoard() model.Board {
	if len(stp) <= 0 {
		return nil
	}
	board := make([]model.Symbols, len(stp[0].Symbols))
	for _, s := range stp {
		for j, currSym := range s.Symbols.ToModelSymbols() {
			board[j] = append(board[j], currSym)
		}
	}
	return board
}

func (ss Symbols) ToModelSymbols() model.Symbols {
	var syms model.Symbols
	for _, s := range ss {
		syms = append(syms, model.Symbol(s))
	}
	return syms
}
