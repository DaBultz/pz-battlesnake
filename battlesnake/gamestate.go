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

// Initialize the game state
func (state *GameState) initialize(options GameOptions) {
	// Generate Game ID
	state.gameID = uuid.New().String()

	// save the options to the state
	state.options = options

	// Load Game Map
	gameMap, err := maps.GetMap(state.options.Map)
	if err != nil {
		log.Fatalf("Error loading map: %v", err)
	}
	state.gameMap = gameMap

	// Create settings Object
	state.settings = map[string]string{
		rules.ParamGameType:            state.options.GameType,
		rules.ParamFoodSpawnChance:     "15",
		rules.ParamMinimumFood:         "1",
		rules.ParamHazardDamagePerTurn: "0",
		rules.ParamShrinkEveryNTurns:   "0",
	}

	// generate seed if not set
	if state.options.Seed == 0 {
		state.options.Seed = int64(time.Now().UTC().UnixNano())
	}

	// Build ruleset from settings
	ruleset := rules.NewRulesetBuilder().
		WithSeed(state.options.Seed).
		WithParams(state.settings).
		WithSolo(len(state.options.Names) == 1).Ruleset()

	state.ruleset = ruleset

	// Create snake states as empty
	state.snakeStates = map[string]SnakeState{}
}

// Build Snake States
func (state *GameState) buildSnakes() map[string]SnakeState {
	bodyChars := []rune{'■', '⌀', '●', '☻', '◘', '☺', '□', '⍟'}
	snakes := map[string]SnakeState{}

	for i := 0; i < len(state.options.Names); i++ {
		name := state.options.Names[i]

		snakes[name] = SnakeState{
			Name: name, ID: name, LastMove: "up",
			Character: bodyChars[i%8],
			Color:     state.options.Colors[i%len(state.options.Colors)],
		}

	}

	return snakes
}

// Build Board State
func (state *GameState) buildBoardState() *rules.BoardState {
	snakeIds := []string{}

	for _, snakeState := range state.snakeStates {
		snakeIds = append(snakeIds, snakeState.ID)
	}

	// Setup Board
	boardState, err := maps.SetupBoard(
		state.gameMap.ID(),
		state.ruleset.Settings(),
		state.options.Height,
		state.options.Width,
		snakeIds,
	)

	if err != nil {
		log.Fatalf("Error Initializing Board State: %v", err)
	}

	// Modify Board State
	boardState, err = state.ruleset.ModifyInitialBoardState(boardState)

	if err != nil {
		log.Fatalf("Error Modifying Board State: %v", err)
	}

	return boardState
}

func (state *GameState) createNextBoardState(boardState *rules.BoardState, moves []rules.SnakeMove) *rules.BoardState {
	// Loop over all snakes and set their last move to the new one
	for _, move := range moves {
		snakeState := state.snakeStates[move.ID]
		snakeState.LastMove = move.Move
		state.snakeStates[move.ID] = snakeState
	}

	boardState, err := state.ruleset.CreateNextBoardState(boardState, moves)

	if err != nil {
		log.Fatalf("Error Creating Next Board State: %v", err)
	}

	boardState, err = maps.UpdateBoard(state.gameMap.ID(), boardState, state.ruleset.Settings())

	if err != nil {
		log.Fatalf("Error Updating Board State: %v", err)
	}

	boardState.Turn += 1

	return boardState
}

func (state *GameState) printMap(boardState *rules.BoardState, useColor bool) {
	var o bytes.Buffer
	o.WriteString(fmt.Sprintf("Ruleset: %s, Seed: %d, Turn: %v\n", state.options.GameType, state.options.Seed, boardState.Turn))
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
		red, green, blue := parseSnakeColor(state.snakeStates[s.ID].Color)
		for _, b := range s.Body {
			if b.X >= 0 && b.X < boardState.Width && b.Y >= 0 && b.Y < boardState.Height {
				if useColor {
					board[b.X][b.Y] = fmt.Sprintf(commands.TERM_FG_RGB+"■", red, green, blue)
				} else {
					board[b.X][b.Y] = string(state.snakeStates[s.ID].Character)
				}
			}
		}
		if useColor {
			o.WriteString(fmt.Sprintf("%v "+commands.TERM_FG_RGB+commands.TERM_BG_WHITE+"■■■"+commands.TERM_RESET+": %v\n", state.snakeStates[s.ID].Name, red, green, blue, s))
		} else {
			o.WriteString(fmt.Sprintf("%v %c: %v\n", state.snakeStates[s.ID].Name, state.snakeStates[s.ID].Character, s))
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

func (state *GameState) getRequestBodyForSnake(boardState *rules.BoardState, snakeState SnakeState) client.SnakeRequest {
	var youSnake rules.Snake

	// Find your snake
	for _, snake := range boardState.Snakes {
		if snake.ID == snakeState.ID {
			youSnake = snake
			break
		}
	}

	return client.SnakeRequest{
		Game:  state.createClientGame(),
		Turn:  boardState.Turn,
		Board: convertStateToBoard(state.boardState, state.snakeStates),
		You:   convertRulesSnake(youSnake, snakeState),
	}
}

func (state *GameState) createClientGame() client.Game {
	return client.Game{
		ID:      state.gameID,
		Timeout: 0,
		Ruleset: client.Ruleset{
			Name:     state.ruleset.Name(),
			Version:  "cli",
			Settings: state.ruleset.Settings(),
		},
		Map: state.gameMap.ID(),
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

func (state *GameState) getResponseForSnake(snake SnakeState) StepRes {
	return StepRes{
		Done:        state.gameOver,
		Reward:      state.getRewardForSnake(snake),
		Info:        nil,
		Observation: state.getRequestBodyForSnake(state.boardState, snake),
	}
}

func (state *GameState) getRewardForSnake(snakeState SnakeState) int {
	return 0
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
