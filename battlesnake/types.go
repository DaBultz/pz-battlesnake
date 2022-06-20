package main

import (
	"github.com/BattlesnakeOfficial/rules/client"
)

type StepRes struct {
	Done        bool                `json:"done"`
	Reward      int                 `json:"reward"`
	Info        interface{}         `json:"info"`
	Observation client.SnakeRequest `json:"observation"`
}

type GameOptions struct {
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Map      string   `json:"map"`
	GameType string   `json:"game_type"`
	Seed     int64    `json:"seed"`
	Names    []string `json:"names"`
	Colors   []string `json:"colors"`
}
