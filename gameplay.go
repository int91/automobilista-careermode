package main

import (
	"fmt"
	"os"
	"os/exec"
)

var SelectedLicense *License

func MainMenuScreen() {
	ClrScreen()
	fmt.Println("Automobilista 2 - Career Mode - By Convulse")
	fmt.Println("ver. 0.0.5")
	fmt.Println("--------------------------------------------")
	fmt.Println("1.) New Game\n2.) Load Game\n3.) Close")
	var a int
	fmt.Scanln(&a)
	if a == 1 {
		SCREEN = NewGameScreen
	} else if a == 2 {
		SCREEN = LoadGameScreen
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
	var in int

	fmt.Println("Enter the name you'd like for your racing team.")
	fmt.Scanln(&teamName)
	fmt.Printf("Is this correct?\n%s\nPress Enter To Accept\n[1] To go back", teamName)
	fmt.Scanln(&in)
	if in == 1 {
		PlayerTeamCreationScreen()
	}

	player.teams = append(player.teams, NewTeam(teamName))
	player.teamName = teamName
}

func GameMenuScreen() {
	ClrScreen()
	var in int

	fmt.Printf("%s Screen", player.teamName)
	if player.teams[0].cars[0] == nil {
		fmt.Printf("1.) Buy Cars\n")
	} else {
		fmt.Printf("1.) Go Race\n")
	}
	fmt.Printf("2.) Take License Test\n")
	fmt.Scanln(&in)
	if in == 1 && player.teams[0].cars[0] == nil {

	} else if in == 1 && player.teams[0].cars[0] != nil {

	} else if in == 2 {

	}

}

func CarMarketScreen() {
	ClrScreen()

	var in int
	fmt.Printf("Vehicle Market")
	fmt.Printf("1.) Browse Vehicles")
	fmt.Printf("2.) Leave")
	fmt.Scanln(&in)

	if in == 1 {
		SCREEN = CarShowroomScreen
	} else if in == 2 {
		SCREEN = GameMenuScreen
	}
}

func CarShowroomScreen() {
	ClrScreen()

	var in int
	//TODO: Only Draw 10 at a time
	for i, ve := range cars {
		fmt.Printf("%d.) %s", i+1, ve.name)
	}
	if in <= len(cars) {
		in -= 1
		cars[in].CarBuyScreen()
	}
}

func (car *Car) CarBuyScreen() {
	ClrScreen()

	var in int
	fmt.Printf("%s\nBuy Price: $%f\nRent Price: $%f\n1.) Purchase\n2.) Rent", car.name, car.price, car.price)
	fmt.Scanln(&in)

	if in == 1 {
		player.BuyCar(car)
	} else if in == 2 {

	} else if in == 3 {
		SCREEN = CarShowroomScreen
	}

}

func RaceSetupScreen() {

}

func RaceScreen() {

}

func LicenseTestSelectScreen() {
	var tests []*License
	for _, lec := range licenses {
		if !player.HasLicense(lec) {
			tests = append(tests, lec)
		}
	}

	if len(tests) > 0 {
		ClrScreen()

		var in int

		for i, license := range tests {
			fmt.Printf("%d.) %s - $%f", i+1, license.name, license.testCost)
		}

		fmt.Scanln(&in)

		if in <= len(tests) {
			in -= 1
			licenses[in].DrawLicenseTestSetupScreen()
		}

	} else {
		SCREEN = GameMenuScreen
	}
}

func LicenseTestScreen() {
	var in int

	fmt.Printf("%s License Test @ %s.\n", SelectedLicense.name, SelectedLicense.track.name)
	fmt.Printf("You have 3 laps to beat a time of %f.\n", SelectedLicense.laptime)
	fmt.Printf("1.) Passed\n")
	fmt.Printf("2.) Failed\n")

	if in == 1 {
		player.driver.licenses = append(player.driver.licenses, SelectedLicense)
	} else if in == 2 {
		SCREEN = GameMenuScreen
	}
}

func NewGameScreen() {
	FirstLaunchScreen()
	PlayerCreationScreen()
	PlayerTeamCreationScreen()
	//TODO: Create Basic Save Game File
	SCREEN = GameMenuScreen
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

func (lic *License) DrawLicenseTestSetupScreen() {
	//TODO: Draw "Press 1 once loaded into the time trial"
	//TODO: If the player does not own the car for the test, they may rent or purchase it.
	ClrScreen()

	var in int
	fmt.Printf("Please load a Practice session on %s with the %s.", lic.track.name, lic.car.name)
	fmt.Printf("You will be given 3 laps to beat a laptime of %f", lic.laptime)
	if !player.HasCar(lic.car.name) {
		fmt.Printf("You currently do not own the %s.\nWould you like to purchase/rent it?\n1.) Car Market\n2.) Back", lic.car.name)
	} else {
		fmt.Printf("1.) Begin Test\n2.) Back")
	}

	if in == 1 && !player.HasCar(lic.car.name) {
		//TODO: Go to car market
		SCREEN = CarMarketScreen
	} else if in == 1 {
		//TODO: Go to license test screen
		if player.HasMoney(false, lic.testCost) {
			player.Pay(false, lic.testCost)
			SelectedLicense = lic
			SCREEN = LicenseTestScreen
		}
	} else if in == 2 {
		SCREEN = LicenseTestSelectScreen
	}
}
