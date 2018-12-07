package machine

import (
	"errors"
	"math/rand"
)

type Position int

type Line struct {
	positions []Position
}

type RandomStopper struct {
	limit int
}

func (s RandomStopper) Stop() int {
	return rand.Intn(s.limit)
}

func NewRandomStopper(limit int) (RandomStopper, error) {
	if limit < 0 {
		return RandomStopper{}, errors.New("Invalid limit")
	}
	return RandomStopper{limit}, nil
}
