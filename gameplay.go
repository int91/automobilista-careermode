package main

import (
	"fmt"
	"os"
	"os/exec"
)

func MainMenuScreen() {
	ClrScreen()
	fmt.Println("Automobilista 2 - Career Mode - By Convulse")
	fmt.Println("ver. 0.0.5")
	fmt.Println("--------------------------------------------")
	fmt.Println("1.) New Game\n2.) Load Game\n3.) Close")
	var a int
	fmt.Scanln(&a)
	if a == 1 {
		NewGameScreen()
	} else if a == 2 {
		LoadGameScreen()
	}
	if a == 3 {
		RUNNING = false
	}
}

func FirstLaunchScreen() {
	ClrScreen()
	if FIRST_LAUNCH {
		fmt.Println("Welcome to the Automobilista 2 Career Mode application.\nKeep in mind that to have the best experience with this addon that you follow every step.\nSteps can be found on screen or in the steps.txt file in the directory.")
		fmt.Println("Press enter to continue")
		var in string
		fmt.Scanln(&in)
	}
}

func PlayerCreationScreen() {
	ClrScreen()
	var nme string
	var num int
	var in int

	fmt.Println("Enter your name")
	fmt.Scanln(&nme)
	fmt.Println("Enter your driver number")
	fmt.Scanln(&num)
	ClrScreen()
	fmt.Printf("Is this information correct?\nName: %s\nDriver Number: %d\nPress Enter To Accept\n[1] To go back", nme, num)
	fmt.Scanln(&in)
	if in == 1 {
		PlayerCreationScreen()
	}

	player = NewPlayer(nme, num)
}

func PlayerTeamCreationScreen() {
	ClrScreen()
	var teamName string

	fmt.Println("Enter the name you'd like for your racing team.")
	fmt.Scanln(&teamName)

	player.teams = append(player.teams, NewTeam(teamName))
}

func NewGameScreen() {
	FirstLaunchScreen()
	PlayerCreationScreen()
	PlayerTeamCreationScreen()
	//TODO: Create Basic Save Game File
	fmt.Println("")
	var in string
	fmt.Scanln(&in)
}

func LoadGameScreen() {

}

func ClrScreen() {
	if RUNTIME == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if RUNTIME == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
