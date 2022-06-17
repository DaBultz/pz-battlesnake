package main

import (
	"C"
	"encoding/json"
	"log"

	"github.com/BattlesnakeOfficial/rules"
)

// encode json from string
func encodeJson(documentPtr *C.char) map[string]interface{} {
	documentString := C.GoString(documentPtr)
	var jsonDocument map[string]interface{}
	err := json.Unmarshal([]byte(documentString), &jsonDocument)

	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	return jsonDocument
}

// get keys as string
func getKeys(jsonDocument map[string]interface{}) []string {
	var keys []string
	for key := range jsonDocument {
		keys = append(keys, key)
	}
	return keys
}

// This convers json to a map of SnakeMoves
func toSnakeMoves(jsonDocument map[string]interface{}) []rules.SnakeMove {
	var snakeMoves []rules.SnakeMove

	for _, key := range getKeys(jsonDocument) {
		snakeMove := rules.SnakeMove{
			ID:   key,
			Move: jsonDocument[key].(string),
		}

		snakeMoves = append(snakeMoves, snakeMove)
	}

	return snakeMoves
}
