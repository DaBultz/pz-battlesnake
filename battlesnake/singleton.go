package main

var game *GameState

// Get Singleton
func GetState() *GameState {
	if game == nil {
		game = &GameState{}
	}
	return game
}
