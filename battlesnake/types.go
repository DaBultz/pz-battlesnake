package main

import (
	"github.com/BattlesnakeOfficial/rules"
	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/BattlesnakeOfficial/rules/maps"
)

type StepRes struct {
	Done        bool                `json:"done"`
	Reward      int                 `json:"reward"`
	Info        interface{}         `json:"info"`
	Observation client.SnakeRequest `json:"observation"`
}

type SnakeState struct {
	Name      string
	ID        string
	LastMove  string
	Character rune
	Color     string
	Head      string
	Tail      string
}

type GameOptions struct {
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Map      string   `json:"map"`
	GameType string   `json:"game_type"`
	Seed     int64    `json:"seed"`
	Names    []string `json:"names"`
}

type GameState struct {
	// Options
	options GameOptions
	// Internal State
	settings    map[string]string
	snakeStates map[string]SnakeState
	gameID      string
	ruleset     rules.Ruleset
	gameMap     maps.GameMap
	boardState  *rules.BoardState
	gameOver    bool
}
