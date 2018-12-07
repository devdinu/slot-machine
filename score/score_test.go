package score

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldComputeScoreForASymbolGivenAPayLine(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[Symbol]symbolScore{
			"sym1": []int{0, 0, 20, 50},
			"sym2": []int{1, 2, 3, 4},
		},
	}
	scorer := Scorer{
		paylines: []Line{[]location{
			{0, 0}, {1, 1}, {2, 2}, {1, 0},
		}},
		card: scorecard,
	}
	board := []Symbols{
		Symbols{"sym1", "sym01", "sym02", "sym03"},
		Symbols{"sym1", "sym1", "sym12", "sym13"},
		Symbols{"sym20", "sym21", "sym22", "sym1"},
	}

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{won: 20}, boardScore)
}

func TestShouldComputeOccurencesUntilSameSymbol(t *testing.T) {
	scorer := Scorer{}
	l := Line{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	testCases := []struct {
		Board
		expectedSym   Symbol
		expectedCount int
	}{
		{
			Board{{"1", "1", "1", "2"}},
			"1", 3,
		},
		{
			Board{{"1", "2", "1", "2"}},
			"1", 1,
		},
		{
			Board{{"0", "1", "1", "2"}},
			"0", 1,
		},
		{
			Board{{"sym", "sym", "sym", "sym"}},
			"sym", 4,
		},
	}

	for _, tc := range testCases {

		occ, err := scorer.findOccurrences(l, tc.Board)

		require.NoError(t, err)
		assert.Equal(t, tc.expectedSym, occ.sym)
		assert.Equal(t, tc.expectedCount, occ.count)
	}
}
