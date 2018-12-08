package game

import (
	"context"
	"testing"

	"github.com/devdinu/slot_machine/machine"
	"github.com/devdinu/slot_machine/score"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestShouldSpinAndComputeScore(t *testing.T) {
	scorer, mach := new(scorerMock), new(machineMock)
	ctx := context.Background()
	r1 := machine.Symbols{"sym1", "sym1", "symo"}
	r2 := machine.Symbols{"symx", "symy", "sym"}
	r3 := machine.Symbols{"sym1", "symbla", "symfoo"}
	stops := []machine.Stop{{Symbols: r1, Position: 1}, {Symbols: r2, Position: 15}, {Symbols: r3, Position: 7}}
	gameScore := score.Score{Won: int64(12345)}
	mach.On("Spin").Return(stops, nil)
	scorer.On("Compute", ctx, Board{r1, r2, r3}).Return(gameScore, nil)

	service := NewService(mach, scorer)
	stopPositions := []int{1, 15, 7}
	bet := int64(2000)

	spin, err := service.SpinOnce(ctx, bet)
	require.NoError(t, err)

	assert.Equal(t, gameScore.Won*bet, spin.Won)
	assert.Equal(t, stopPositions, spin.Stops)
}

type machineMock struct{ mock.Mock }

func (mm *machineMock) Spin() ([]machine.Stop, error) {
	args := mm.Called()
	return args.Get(0).([]machine.Stop), args.Error(1)
}

type scorerMock struct{ mock.Mock }

func (sm *scorerMock) Compute(ctx context.Context, board Board) (score.Score, error) {
	args := sm.Called(ctx, board)
	return args.Get(0).(score.Score), args.Error(1)
}
