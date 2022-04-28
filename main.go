package main

import "runtime"

var licenses []*License
var drivers []*Driver
var cars []*Car
var tracks []*Track

var player *Player
var FIRST_LAUNCH bool
var RUNNING bool
var RUNTIME string
var SCREEN func()

func main() {
	RUNTIME = runtime.GOOS
	RUNNING = true
	SCREEN = MainMenuScreen
	//TODO: Load First Launch file (if it exists get the value)
	FIRST_LAUNCH = true
	SetupTest()

	for RUNNING {
		SCREEN()
	}
}
