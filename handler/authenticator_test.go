package handler

import (
	"testing"

	"github.com/devdinu/slot_machine/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticate(t *testing.T) {
	config.Load("../config.yaml")
	tokens := []string{
		`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidWlkIjoidXNlci1pZCIsImNoaXBzIjoxMjM0NSwiYmV0Ijo1MDAsImlhdCI6MTUxNjIzOTAyMn0.zeT63Njv3blEMhYd8WCW6P63wcnoFTTKFoQiNX_yN70`,
		`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidWlkIjoidXNlci1pZCIsImNoaXBzIjoxMjM0NSwiYmV0Ijo1MDAsImlhdCI6MTU0NDI0ODk3NywiZXhwIjoyODA2NTUyOTc3fQ.r5o41ZBolMHc4aMOrkH_1x_w6zq1FMX0jW1_vyEEsqw`,
	}

	for _, tok := range tokens {
		claim, err := authenticate(tok)

		require.NoError(t, err, tok)
		assert.Equal(t, claim.Bet, int64(500))
		assert.Equal(t, claim.Chips, int64(12345))
		assert.Equal(t, claim.UID, "user-id")
	}

}
