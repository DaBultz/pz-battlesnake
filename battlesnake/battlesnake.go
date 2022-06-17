package main

import (
	"C"
	"log"
	"time"

	"github.com/BattlesnakeOfficial/rules"
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
	Width  int
	Height int
	Seed   int64
	Names  []string

	// Internal State
	gameNum     int
	settings    map[string]string
	snakeStates map[string]SnakeState
	gameID      string
	ruleset     rules.Ruleset
	gameMap     maps.GameMap
	boardState  *rules.BoardState
	gameOver    bool
}

// Initialize the game state
func (gameState *GameState) initialize() {
	// Generate Game ID
	gameState.gameID = uuid.New().String()

	// Load Game Map
	gameMap, err := maps.GetMap("standard")
	if err != nil {
		log.Fatalf("Error loading map: %v", err)
	}
	gameState.gameMap = gameMap

	// Create settings Object
	gameState.settings = map[string]string{
		rules.ParamGameType:            "solo",
		rules.ParamFoodSpawnChance:     "15",
		rules.ParamMinimumFood:         "1",
		rules.ParamHazardDamagePerTurn: "0",
		rules.ParamShrinkEveryNTurns:   "0",
	}

	// generate seed if not set
	if gameState.Seed == 0 {
		gameState.Seed = int64(time.Now().UTC().UnixNano())
	}

	// Build ruleset from settings
	ruleset := rules.NewRulesetBuilder().
		WithSeed(gameState.Seed).
		WithParams(gameState.settings).
		WithSolo(true).Ruleset()

	gameState.ruleset = ruleset

	// Create snake states as empty
	gameState.snakeStates = map[string]SnakeState{}
}

// Build Snake States
func (gameState *GameState) buildSnakes() map[string]SnakeState {
	bodyChars := []rune{'■', '⌀', '●', '☻', '◘', '☺', '□', '⍟'}
	snakes := map[string]SnakeState{}

	for i := 0; i < len(gameState.Names); i++ {
		name := gameState.Names[i]

		snakes[name] = SnakeState{
			Name: name, ID: name, LastMove: "up", Character: bodyChars[i%8],
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
		gameState.Width,
		gameState.Height,
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
		Game:  game.createClientGame(),
		Turn:  boardState.Turn,
		Board: convertStateToBoard(boardState, gameState.snakeStates),
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

func (gameState *GameState) getRewardForSnake() int {
	snake := gameState.boardState.Snakes[0]

	// Check if snake fills the board
	if len(snake.Body) == (gameState.Width * gameState.Height) {
		return 1
	}

	// check if snake is not dead
	if snake.EliminatedCause != rules.NotEliminated {
		return -1
	}

	return 0
}
