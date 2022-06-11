package main

import (
	"C"
	"encoding/json"
	"fmt"

	"github.com/BattlesnakeOfficial/rules"
)

//export fromJSON
func fromJSON(documentPtr *C.char) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(C.GoString(documentPtr)), &document)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(document)
}

//export setup
func setup() {

}

func main() {
	snake := rules.BoardState{}
	fmt.Println(snake)
}

// Step 1: Create a new GameState
// Step 2: Initialize the GameState
// Step 3: call step() from python
// Step 4:
