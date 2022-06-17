package main

import (
	"C"
	"fmt"
	"math/rand"
)
import (
	"encoding/json"

	"github.com/BattlesnakeOfficial/rules/client"
)

type StepRes struct {
	Done        bool                `json:"done"`
	Reward      int                 `json:"reward"`
	Info        interface{}         `json:"info"`
	Observation client.SnakeRequest `json:"observation"`
}

//export setup
func setup() {
	state := GetState()

	// TODO: Make this configurable from python
	// Set Default Width and Height
	state.Width = 11
	state.Height = 11
	// Set a Default snake
	state.Names = append(state.Names, "agent_1")

	// Initialize the game state
	state.initialize()

	// Setup Snakes
	state.snakeStates = state.buildSnakes()

	// Randmomize the seed
	rand.Seed(state.Seed)

	boardState := state.buildBoardState()
	state.boardState = boardState
}

//export reset
func reset() {
	resetSingleton()
	setup()
}

//export step
func step(actions *C.char) *C.char {
	state := GetState()

	// Check if the game is over advice to reset the environment
	if state.gameOver {
		fmt.Println("Game is over, please call reset()")
		return C.CString("")
	}

	// Convert String Encoded JSON into a map
	moves := toSnakeMoves(encodeJson(actions))

	// Update the board state
	state.boardState = state.createNextBoardState(state.boardState, moves)

	// Check if the game is over
	isGameOver, _ := state.ruleset.IsGameOver(state.boardState)

	if isGameOver {
		state.gameOver = true
	}

	//
	// Build the response
	response := &StepRes{
		Done:        state.gameOver,
		Reward:      state.getRewardForSnake(),
		Info:        nil,
		Observation: state.getRequestBodyForSnake(state.boardState, state.snakeStates[state.Names[0]]),
	}

	// Convert the response to a json
	responseString, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	return C.CString(string(responseString))
}

// empty main
func main() {}
