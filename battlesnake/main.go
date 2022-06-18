package main

import (
	"C"
)
import (
	"encoding/json"
	"fmt"
	"math/rand"
)

//export setup
func setup(settingsDoc *C.char) {
	// This will reset the state if setup is called while a state exists
	// TODO: Find a better way of resetting the enviorment (NEEDED FOR VECTORIZATION OF ENVIORMENT)
	if game != nil {
		game = nil
	}

	// Get the state
	state := GetState()

	// Convert String Encoded JSON into a map
	gameOptions := parseGameOptions(settingsDoc)

	// Initialize the game state
	state.initialize(gameOptions)

	// Build Snake States
	state.snakeStates = state.buildSnakes()

	// Randomize the seed
	rand.Seed(state.options.Seed)

	// Build Boardstate
	state.boardState = state.buildBoardState()
}

//export isGameOver
func isGameOver() C.int {
	state := GetState()

	if state.gameOver {
		return C.int(1)
	}

	return C.int(0)
}

//export step
func step(actions *C.char) *C.char {
	state := GetState()

	// conver the string to a map over snake moves
	moves := parseSnakeMoves(encodeJson(actions))

	// Update the board state
	state.boardState = state.createNextBoardState(state.boardState, moves)

	// Check if the game is over
	done, _ := state.ruleset.IsGameOver(state.boardState)
	if done {
		state.gameOver = true
	}

	// Construct the step response
	response := map[string]StepRes{}

	for _, snakeState := range state.snakeStates {
		response[snakeState.ID] = state.getResponseForSnake(snakeState)
	}

	responseString, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}

	return C.CString(string(responseString))
}

// empty main
func main() {}
