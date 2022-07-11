/**
  Copyright © 2022 Battlesnake Inc.

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/BattlesnakeOfficial/rules/cli/commands"
	"github.com/BattlesnakeOfficial/rules/client"
	"github.com/BattlesnakeOfficial/rules/maps"
	"github.com/google/uuid"
)

type SnakeState struct {
	Name      string
	ID        string
	LastMove  string
	Character rune
	Color     string
	Head      string
	Tail      string
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

// Initialize the game state
func (gameState *GameState) initialize(options GameOptions) {
	// Generate Game ID
	gameState.gameID = uuid.New().String()

	// save the options to the state
	gameState.options = options

	// Load Game Map
	gameMap, err := maps.GetMap(gameState.options.Map)
	if err != nil {
		log.Fatalf("Error loading map: %v", err)
	}
	gameState.gameMap = gameMap

	// Create settings Object
	gameState.settings = map[string]string{
		rules.ParamGameType:            gameState.options.GameType,
		rules.ParamFoodSpawnChance:     "15",
		rules.ParamMinimumFood:         "1",
		rules.ParamHazardDamagePerTurn: "0",
		rules.ParamShrinkEveryNTurns:   "0",
	}

	// generate seed if not set
	if gameState.options.Seed == 0 {
		gameState.options.Seed = int64(time.Now().UTC().UnixNano())
	}

	// Build ruleset from settings
	ruleset := rules.NewRulesetBuilder().
		WithSeed(gameState.options.Seed).
		WithParams(gameState.settings).
		WithSolo(len(gameState.options.Names) == 1).Ruleset()

	gameState.ruleset = ruleset

	// Create snake states as empty
	gameState.snakeStates = map[string]SnakeState{}
}

// Build Snake States
func (gameState *GameState) buildSnakes() map[string]SnakeState {
	bodyChars := []rune{'■', '⌀', '●', '☻', '◘', '☺', '□', '⍟'}
	snakes := map[string]SnakeState{}

	for i := 0; i < len(gameState.options.Names); i++ {
		name := gameState.options.Names[i]

		snakes[name] = SnakeState{
			Name: name, ID: name, LastMove: "up",
			Character: bodyChars[i%8],
			Color:     gameState.options.Colors[i%len(gameState.options.Colors)],
		}

	}

	return snakes
}

// Build Board State
func (gameState *GameState) buildBoardState() *rules.BoardState {
	snakeIds := []string{}

	for _, snakeState := range gameState.snakeStates {
		snakeIds = append(snakeIds, snakeState.ID)
	}

	// Setup Board
	boardState, err := maps.SetupBoard(
		gameState.gameMap.ID(),
		gameState.ruleset.Settings(),
		gameState.options.Height,
		gameState.options.Width,
		snakeIds,
	)

	if err != nil {
		log.Fatalf("Error Initializing Board State: %v", err)
	}

	// Modify Board State
	boardState, err = gameState.ruleset.ModifyInitialBoardState(boardState)

	if err != nil {
		log.Fatalf("Error Modifying Board State: %v", err)
	}

	return boardState
}

func (gameState *GameState) createNextBoardState(boardState *rules.BoardState, moves []rules.SnakeMove) *rules.BoardState {
	// Loop over all snakes and set their last move to the new one
	for _, move := range moves {
		snakeState := gameState.snakeStates[move.ID]
		snakeState.LastMove = move.Move
		gameState.snakeStates[move.ID] = snakeState
	}

	boardState, err := gameState.ruleset.CreateNextBoardState(boardState, moves)

	if err != nil {
		log.Fatalf("Error Creating Next Board State: %v", err)
	}

	boardState, err = maps.UpdateBoard(gameState.gameMap.ID(), boardState, gameState.ruleset.Settings())

	if err != nil {
		log.Fatalf("Error Updating Board State: %v", err)
	}

	boardState.Turn += 1

	return boardState
}

func (gameState *GameState) printMap(boardState *rules.BoardState, useColor bool) {
	var o bytes.Buffer
	o.WriteString(fmt.Sprintf("Ruleset: %s, Seed: %d, Turn: %v\n", gameState.options.GameType, gameState.options.Seed, boardState.Turn))
	board := make([][]string, boardState.Width)

	for i := range board {
		board[i] = make([]string, boardState.Height)
	}

	// Add all empty spaces to the buffer
	for y := int(0); y < boardState.Height; y++ {
		for x := int(0); x < boardState.Width; x++ {
			if useColor {
				board[x][y] = commands.TERM_FG_LIGHTGRAY + "□"
			} else {
				board[x][y] = "◦"
			}

		}
	}

	// Add all hazard to the buffer
	for _, oob := range boardState.Hazards {
		if useColor {
			board[oob.X][oob.Y] = commands.TERM_BG_GRAY + " " + commands.TERM_BG_WHITE
		} else {
			board[oob.X][oob.Y] = "░"
		}

	}

	if useColor {
		o.WriteString(fmt.Sprintf("Hazards "+commands.TERM_BG_GRAY+" "+commands.TERM_RESET+": %v\n", boardState.Hazards))
	} else {
		o.WriteString(fmt.Sprintf("Hazards ░: %v\n", boardState.Hazards))
	}

	// Add all food to the buffer
	for _, f := range boardState.Food {
		if useColor {
			board[f.X][f.Y] = commands.TERM_FG_FOOD + "●"
		} else {
			board[f.X][f.Y] = "⚕"
		}
	}

	if useColor {
		o.WriteString(fmt.Sprintf("Food "+commands.TERM_FG_FOOD+commands.TERM_BG_WHITE+"●"+commands.TERM_RESET+": %v\n", boardState.Food))
	} else {
		o.WriteString(fmt.Sprintf("Food ⚕: %v\n", boardState.Food))
	}

	// Add all snakes to the buffer
	for _, s := range boardState.Snakes {
		// FIXES: https://github.com/DaBultz/pz-battlesnake/issues/10
		if s.EliminatedCause == rules.NotEliminated {
			continue
		}

		red, green, blue := parseSnakeColor(gameState.snakeStates[s.ID].Color)
		for _, b := range s.Body {
			if b.X >= 0 && b.X < boardState.Width && b.Y >= 0 && b.Y < boardState.Height {
				if useColor {
					board[b.X][b.Y] = fmt.Sprintf(commands.TERM_FG_RGB+"■", red, green, blue)
				} else {
					board[b.X][b.Y] = string(gameState.snakeStates[s.ID].Character)
				}
			}
		}

		if useColor {
			o.WriteString(fmt.Sprintf("%v "+commands.TERM_FG_RGB+commands.TERM_BG_WHITE+"■■■"+commands.TERM_RESET+": %v\n", gameState.snakeStates[s.ID].Name, red, green, blue, s))
		} else {
			o.WriteString(fmt.Sprintf("%v %c: %v\n", gameState.snakeStates[s.ID].Name, gameState.snakeStates[s.ID].Character, s))
		}
	}

	// Add border to the buffer
	for y := boardState.Height - 1; y >= 0; y-- {
		if useColor {
			o.WriteString(commands.TERM_BG_WHITE)
		}
		for x := int(0); x < boardState.Width; x++ {
			o.WriteString(board[x][y])
		}
		if useColor {
			o.WriteString(commands.TERM_RESET)
		}
		o.WriteString("\n")
	}

	log.Print(o.String())
}

func (gameState *GameState) getRequestBodyForSnake(boardState *rules.BoardState, snakeState SnakeState) client.SnakeRequest {
	var youSnake rules.Snake

	// Find your snake
	for _, snake := range boardState.Snakes {
		if snake.ID == snakeState.ID {
			youSnake = snake
			break
		}
	}

	return client.SnakeRequest{
		Game:  gameState.createClientGame(),
		Turn:  boardState.Turn,
		Board: convertStateToBoard(gameState.boardState, gameState.snakeStates),
		You:   convertRulesSnake(youSnake, snakeState),
	}
}

func (gameState *GameState) createClientGame() client.Game {
	return client.Game{
		ID:      gameState.gameID,
		Timeout: 0,
		Ruleset: client.Ruleset{
			Name:     gameState.ruleset.Name(),
			Version:  "cli",
			Settings: gameState.ruleset.Settings(),
		},
		Map: gameState.gameMap.ID(),
	}
}

func convertStateToBoard(boardState *rules.BoardState, snakeStates map[string]SnakeState) client.Board {
	return client.Board{
		Height:  boardState.Height,
		Width:   boardState.Width,
		Food:    client.CoordFromPointArray(boardState.Food),
		Hazards: client.CoordFromPointArray(boardState.Hazards),
		Snakes:  convertRulesSnakes(boardState.Snakes, snakeStates),
	}
}

func convertRulesSnake(snake rules.Snake, snakeState SnakeState) client.Snake {
	return client.Snake{
		ID:      snake.ID,
		Name:    snakeState.Name,
		Health:  snake.Health,
		Body:    client.CoordFromPointArray(snake.Body),
		Latency: "0",
		Head:    client.CoordFromPoint(snake.Body[0]),
		Length:  int(len(snake.Body)),
		Shout:   "",
		Customizations: client.Customizations{
			Head:  snakeState.Head,
			Tail:  snakeState.Tail,
			Color: snakeState.Color,
		},
	}
}

func convertRulesSnakes(snakes []rules.Snake, snakeStates map[string]SnakeState) []client.Snake {
	a := make([]client.Snake, 0)
	for _, snake := range snakes {
		if snake.EliminatedCause == rules.NotEliminated {
			a = append(a, convertRulesSnake(snake, snakeStates[snake.ID]))
		}
	}
	return a
}

func (gameState *GameState) getRewardForSnake(snakeState SnakeState) int {
	// Find Snake index
	var youSnake rules.Snake

	for _, snake := range gameState.boardState.Snakes {
		if snake.ID == snakeState.ID {
			youSnake = snake
			break
		}
	}

	// Game is in progress
	if !gameState.gameOver {
		return 0
	}

	if youSnake.EliminatedCause != rules.NotEliminated {
		return -1
	}

	return 1
}

func (gameState *GameState) isSnakeDone(snakeState SnakeState) bool {
	// Find Snake index
	var youSnake rules.Snake

	for _, snake := range gameState.boardState.Snakes {
		if snake.ID == snakeState.ID {
			youSnake = snake
			break
		}
	}

	return youSnake.EliminatedCause != rules.NotEliminated
}

func (gameState *GameState) getResponseForSnake(snake SnakeState) StepRes {
	return StepRes{
		Done:        gameState.isSnakeDone(snake),
		Reward:      gameState.getRewardForSnake(snake),
		Info:        nil,
		Observation: gameState.getRequestBodyForSnake(gameState.boardState, snake),
	}
}

// Parses a color string like "#ef03d3" to rgb values from 0 to 255 or returns
// the default gray if any errors occure
func parseSnakeColor(color string) (int64, int64, int64) {
	if len(color) == 7 {
		red, err_r := strconv.ParseInt(color[1:3], 16, 64)
		green, err_g := strconv.ParseInt(color[3:5], 16, 64)
		blue, err_b := strconv.ParseInt(color[5:], 16, 64)
		if err_r == nil && err_g == nil && err_b == nil {
			return red, green, blue
		}
	}
	// Default gray color from Battlesnake board
	return 136, 136, 136
}
