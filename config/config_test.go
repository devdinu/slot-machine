package config

import (
	"testing"
	"time"

	model "github.com/devdinu/slot_machine/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigLoad(t *testing.T) {
	filename := "../config.yaml"
	err := Load(filename)
	require.NoError(t, err)

	assert.Equal(t, []byte("ThisIsTopSecret"), AuthSecret())
	assert.True(t, AuthEnabled())
	assert.Equal(t, time.Duration(30)*time.Minute, AuthTokenExpiryMinutes())

	gameCfg := Gaming()
	assert.Equal(t, 5, len(gameCfg.ReelsOfSymbols))
	assert.Equal(t, 3, gameCfg.Rows)
	assert.Equal(t, "Atkins", gameCfg.Wild.String())
	assert.Equal(t, model.Symbol("Scale"), gameCfg.Scatter)

	score := Scoring()
	assert.Equal(t, 4, len(score.Paylines))
	assert.Equal(t, 10, len(score.SymbolsScore))
}
