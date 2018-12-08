package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigLoad(t *testing.T) {
	filename := "../config.yaml"
	err := Load(filename)
	require.NoError(t, err)

	assert.Equal(t, []byte("ThisIsTopSecret"), AuthSecret())
	assert.True(t, AuthEnabled())

	gameCfg := Gaming()
	assert.Equal(t, 5, len(gameCfg.Reels))
	assert.Equal(t, 3, gameCfg.Rows)

	score := Scoring()
	assert.Equal(t, 3, len(score.Paylines))
	assert.Equal(t, 10, len(score.SymbolsScore))
}
