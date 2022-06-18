package main

import (
	"log"
	"time"

	"github.com/BattlesnakeOfficial/rules"
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
			Name: name, ID: name, LastMove: "up", Character: bodyChars[i%8],
		}

	}

	return snakes
}

// Build Board State
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
		Reward:      0,
		Info:        nil,
		Observation: state.getRequestBodyForSnake(state.boardState, snake),
	}
}
