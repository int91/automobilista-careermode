package main

import "runtime"

var licenses []*License
var drivers []*Driver
var cars []*Car
var player *Player
var FIRST_LAUNCH bool
var RUNNING bool
var RUNTIME string

func main() {
	RUNTIME = runtime.GOOS
	RUNNING = true
	//TODO: Load First Launch file (if it exists get the value)
	FIRST_LAUNCH = true
	SetupTest()
	for RUNNING {
		MainMenuScreen()
	}
}
