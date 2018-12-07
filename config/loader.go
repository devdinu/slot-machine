package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Application struct {
	Name string
	Score
	Game
	Server
}

type Score struct {
	SymbolsScore map[string][]int `yaml:"symbol_score"`
	Paylines     []Line           `yaml:"pay_lines,flow"`
}

type Line []Location

type Server struct {
	Port int
}

type Location struct {
	Row int
	Col int
}

type Game struct {
	Reels []Reel
	Rows  int
}

type Reel []string

var appConfig Application

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
