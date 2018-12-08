package score

import (
	"context"
	"testing"

	model "github.com/devdinu/slot_machine/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldComputeScoreWithWildSymbol(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[model.Symbol]symbolScore{
			"sym1":    []int{0, 10, 20, 50},
			"scatter": []int{0, 100, 200, 300},
		},
	}
	scorer := Scorer{
		paylines: []model.Line{
			{{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}, {Row: 0, Col: 3}}, // sym1:4 = 50
		},
		card:    scorecard,
		scatter: "scatter",
	}
	board := []model.Symbols{
		{"sym1", "sym1", "sym2", "scatter"},
		{"sym2", "sym1", "scatter", "sym13"},
		{"sym3", "scatter", "sym3", "sym1"},
		{"sym3", "sym3", "sym3", "sym1"},
	}
	expectedScore := int64(10 + 200)

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{expectedScore}, boardScore, "Should add scatter score too")
}

func TestShouldConsiderWildcardSymbol(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[model.Symbol]symbolScore{
			"sym1": []int{0, 10, 30, 50},
			"sym2": []int{0, 20, 20, 50},
			"wild": []int{0, 100, 200, 300},
			"sym3": []int{0, 100, 200, 300},
			"sym4": []int{0, 100, 200, 400},
		},
	}
	scorer := Scorer{
		paylines: []model.Line{
			{{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}, {Row: 0, Col: 3}}, // row1
			{{Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 1, Col: 2}, {Row: 1, Col: 3}}, // row2
			{{Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2}, {Row: 2, Col: 3}}, // row3
			{{Row: 3, Col: 0}, {Row: 3, Col: 1}, {Row: 3, Col: 2}, {Row: 3, Col: 3}}, // row4
		},
		card: scorecard,
		wild: "wild",
	}
	board := []model.Symbols{
		{"sym1", "sym1", "wild", "symx"}, // thrice: 30
		{"sym2", "wild", "bla", "sym13"}, // twice: 20
		{"sym3", "wild", "sym3", "sym1"}, // thrice: 200
		{"sym4", "wild", "wild", "sym4"}, // four: 400
	}
	expectedScore := int64(30 + 20 + 200 + 400)

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{expectedScore}, boardScore, "Should consider wildcard symbol too")

}

func TestShouldComputeScoreForMultiplePayLines(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[model.Symbol]symbolScore{
			"sym1":    []int{0, 0, 20, 50},
			"sym2":    []int{1, 2, 3, 4},
			"sym3":    []int{100, 200, 300, 400},
			"unknown": []int{100, 200, 300, 400},
		},
	}
	scorer := Scorer{
		paylines: []model.Line{
			{{Row: 0, Col: 0}, {Row: 1, Col: 1}, {Row: 1, Col: 2}, {Row: 2, Col: 3}}, // sym1:4 = 50
			{{Row: 1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}, {Row: 1, Col: 3}}, // sym2:3 = 3
			{{Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2}, {Row: 0, Col: 3}}, // sym3:4: 400
		},
		card: scorecard,
	}
	board := []model.Symbols{
		{"sym1", "sym2", "sym2", "sym3"},
		{"sym2", "sym1", "sym1", "sym13"},
		{"sym3", "sym3", "sym3", "sym1"},
		{"sym3", "sym3", "sym3", "sym1"},
	}
	expectedScore := int64(50 + 3 + 400)

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{expectedScore}, boardScore)
}

func TestShouldComputeScoreForASymbolGivenAPayLine(t *testing.T) {
	scorecard := scoreCard{
		symbolScores: map[model.Symbol]symbolScore{
			"sym2": []int{1, 2, 3, 4},
			"sym1": []int{0, 10, 20, 50},
		},
	}
	scorer := Scorer{
		paylines: []model.Line{[]model.Location{
			{Row: 0, Col: 0}, {Row: 1, Col: 1}, {Row: 2, Col: 0}, {Row: 1, Col: 1},
		}},
		card: scorecard,
	}
	board := []model.Symbols{
		{"sym1", "sym01", "sym1", "sym03"},
		{"sym1", "sym1", "sym12", "sym13"},
		{"sym20", "sym21", "sym22", "sym1"},
	}

	boardScore, err := scorer.Compute(context.Background(), board)

	require.NoError(t, err)
	assert.Equal(t, Score{10}, boardScore)
}

func TestShouldComputeOccurencesUntilSameSymbol(t *testing.T) {
	scorer := Scorer{}
	l := model.Line{{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}, {Row: 0, Col: 3}}
	testCases := []struct {
		model.Board
		expectedSym   model.Symbol
		expectedCount int
	}{
		{
			model.Board{{"1", "1", "1", "2"}},
			"1", 3,
		},
		{
			model.Board{{"1", "2", "1", "2"}},
			"1", 1,
		},
		{
			model.Board{{"0", "1", "1", "2"}},
			"0", 1,
		},
		{
			model.Board{{"sym", "sym", "sym", "sym"}},
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
