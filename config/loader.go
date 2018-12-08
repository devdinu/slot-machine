package config

import (
	"os"

	"github.com/devdinu/slot_machine/machine"
	"gopkg.in/yaml.v2"
)

var appConfig Application

type Line []Location
type Server struct{ Port int }

type Application struct {
	Name string
	Score
	Game
	Server
	Authentication
}
type Score struct {
	SymbolsScore map[string][]int `yaml:"symbol_score"`
	Paylines     []Line           `yaml:"pay_lines,flow"`
}
type Location struct {
	Row int
	Col int
}
type Game struct {
	ReelsOfSymbols []machine.Symbols `yaml:"reels"`
	Rows           int
	Scatter        []machine.Symbol `yaml:"scatter"`
	Wild           []machine.Symbol `yaml:"wild"`
}

func Load(file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}

	err = yaml.NewDecoder(reader).Decode(&appConfig)
	if err != nil {
		return err
	}
	//TODO: validate configuration
	return nil
}
