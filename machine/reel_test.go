package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnSymbolsForGivenPosition(t *testing.T) {
	stopper := new(stopperMock)
	stopper.On("Stop").Return(Position(1))
	reel := Reel{
		[]Symbol{"sym1", "sym2", "sym3", "sym4"},
		stopper,
		2,
	}

	stop := reel.Spin()

	require.Equal(t, 2, len(stop.Symbols))
	assert.Equal(t, Symbol("sym2"), stop.Symbols[0])
	assert.Equal(t, Symbol("sym3"), stop.Symbols[1])
}

func TestShouldReturnSymbolsForGivenPositionInCycle(t *testing.T) {
	stopper := new(stopperMock)
	stopper.On("Stop").Return(Position(2))
	reel := Reel{
		[]Symbol{"sym1", "sym2", "sym3", "sym4"},
		stopper,
		3,
	}

	stop := reel.Spin()

	require.Equal(t, 3, len(stop.Symbols))
	assert.Equal(t, Symbols{"sym3", "sym4", "sym1"}, stop.Symbols)
	assert.Equal(t, Position(2), stop.Position)
}
