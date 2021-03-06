package model

type Symbol string
type Symbols []Symbol
type Board []Symbols

func (s Symbol) String() string {
	return string(s)
}

func (b Board) Get(l Location) Symbol {
	return b[l.Row][l.Col]
}

func (b Board) Empty() bool {
	return len(b) == 0 || len(b[0]) == 0
}

type Line []Location

type Location struct {
	Row int
	Col int
}

type Position int

type CtxKey string

var UserBetKey = CtxKey("bet")
var UserIDKey = CtxKey("uid")
var UserChipsKey = CtxKey("chips")
