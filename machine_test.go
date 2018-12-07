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
	)
	stopper.On("Stop").Return(Position(1))
	expectedStop := Stop{chosen: []Symbol{"butter"}}

	stops, err := machine.Spin()

	require.NoError(t, err)
	require.Equal(t, 1, len(stops))
	require.Equal(t, 1, len(stops[0].chosen))
	assert.Equal(t, expectedStop, stops[0])
	stopper.AssertExpectations(t)
}

type stopperMock struct{ mock.Mock }

func (s *stopperMock) Stop() Position {
	return s.Called().Get(0).(Position)
}
