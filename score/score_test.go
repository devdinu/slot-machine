package score

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldComputeScoreForMultiplePayLines(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[Symbol]symbolScore{
			"sym1":    []int{0, 0, 20, 50},
			"sym2":    []int{1, 2, 3, 4},
			"sym3":    []int{100, 200, 300, 400},
			"unknown": []int{100, 200, 300, 400},
		},
	}
	scorer := Scorer{
		paylines: []Line{
			{{0, 0}, {1, 1}, {1, 2}, {2, 3}}, // sym1:4 = 50
			{{1, 0}, {0, 1}, {0, 2}, {1, 3}}, // sym2:3 = 3
			{{2, 0}, {2, 1}, {2, 2}, {0, 3}}, // sym3:4: 400
		},
		card: scorecard,
	}
	board := []Symbols{
		Symbols{"sym1", "sym2", "sym2", "sym3"},
		Symbols{"sym2", "sym1", "sym1", "sym13"},
		Symbols{"sym3", "sym3", "sym3", "sym1"},
	}
	expectedScore := int64(50 + 3 + 400)

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{expectedScore}, boardScore)
}

func TestShouldComputeScoreForASymbolGivenAPayLine(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[Symbol]symbolScore{
			"sym2": []int{1, 2, 3, 4},
			"sym1": []int{0, 10, 20, 50},
		},
	}
	scorer := Scorer{
		paylines: []Line{[]location{
			{0, 0}, {1, 1}, {2, 0}, {1, 1},
		}},
		card: scorecard,
	}
	board := []Symbols{
		Symbols{"sym1", "sym01", "sym1", "sym03"},
		Symbols{"sym1", "sym1", "sym12", "sym13"},
		Symbols{"sym20", "sym21", "sym22", "sym1"},
	}

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{10}, boardScore)
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
