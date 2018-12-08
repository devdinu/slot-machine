package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestShouldReturnStoppedSymbolInReel(t *testing.T) {
	stopper := new(stopperMock)
	machine := NewMachine(
		stopper,
		[]Symbol{"cheese", "butter", "eggs"},
		1,
		1,
	)
	stopPosition := Position(1)
	stopper.On("Stop").Return(stopPosition)
	expectedStop := Stop{[]Symbol{"butter"}, stopPosition}

	stops, err := machine.Spin()

	require.NoError(t, err)
	require.Equal(t, 1, len(stops))
	require.Equal(t, 1, len(stops[0].Symbols))
	assert.Equal(t, expectedStop, stops[0])
	stopper.AssertExpectations(t)
}

func TestShouldReturnStoppedSymbolsForReels(t *testing.T) {
	stopper := new(stopperMock)
	machine := NewMachine(
		stopper,
		[]Symbol{"cheese", "butter", "eggs", "something", "else"},
		2,
		2,
	)
	stopper.On("Stop").Return(Position(0)).Once()
	stopper.On("Stop").Return(Position(2)).Once()
	expectedStop1 := Stop{[]Symbol{"cheese", "butter"}, 0}
	expectedStop2 := Stop{[]Symbol{"eggs", "something"}, 2}

	stops, err := machine.Spin()

	require.NoError(t, err)
	require.Equal(t, 2, len(stops))
	assert.Equal(t, expectedStop1, stops[0])
	assert.Equal(t, expectedStop2, stops[1])
	stopper.AssertExpectations(t)
}

type stopperMock struct{ mock.Mock }

func (s *stopperMock) Stop() Position {
	return s.Called().Get(0).(Position)
}
