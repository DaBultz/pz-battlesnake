package main

import (
	"github.com/BattlesnakeOfficial/rules"
	"github.com/BattlesnakeOfficial/rules/cli/commands"
	"github.com/BattlesnakeOfficial/rules/maps"
)

type GameState struct {
	Width  int
	Height int

	// Internal state
	settings    map[string]string
	snakeStates map[string]commands.SnakeState
	gameID      string
	ruleset     rules.Ruleset
	gameMap     maps.GameMap
}
