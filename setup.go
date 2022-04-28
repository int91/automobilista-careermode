package main

import (
	"fmt"
)

var NATIONS []string

var FULT_ENTRY_FEE float64 = 899999.99
var KART125_ENTRY_FEE float64 = 29999.99
var KARTSHIFTER_ENTRY_FEE float64 = 31999.99
var SUPERKART_ENTRY_FEE float64 = 40999.99

func CreateTracks() {
	AddTrack(NewTrack("Interlagos GP", 4.29, 15, 89999.99, false))
	AddTrack(NewTrack("Imola", 4.9, 19, 115999.99, false))
	AddTrack(NewTrack("Hockenheim", 4.57, 17, 63996.99, true))
	AddTrack(NewTrack("Brands Hatch", 3.91, 10, 49833.89, false))
	AddTrack(NewTrack("Monaco", 3.33, 19, 99999.99, false))
	AddTrack(NewTrack("Leguna Seca", 3.6, 11, 39999.99, true))
	AddTrack(NewTrack("Long Beach", 3.16, 11, 64589.99, true))
	AddTrack(NewTrack("Montreal", 4.36, 14, 86999.99, false))
	AddTrack(NewTrack("Monza", 5.79, 11, 130999.99, true))
	AddTrack(NewTrack("Nurburgring GP", 5.14, 15, 101999.99, true))
	AddTrack(NewTrack("Road America", 6.51, 14, 178989.59, true))
	AddTrack(NewTrack("SpielBerg", 4.31, 10, 59999.99, false))
	AddTrack(NewTrack("Santa Cruz do Sul", 3.32, 14, 39999.99, false))
	AddTrack(NewTrack("Silverstone", 5.89, 18, 149593.99, true))
	AddTrack(NewTrack("Curvelo Long", 4.42, 18, 109999.99, false))
	AddTrack(NewTrack("Cascais Alternate", 4.15, 11, 99519.99, false))
	AddTrack(NewTrack("Spa", 0, 0, 0, true)) // ADD Data
}

func CreateLicenses() {
	AddLicense(NewLicense("Rookie", 0))
	AddLicense(NewLicense("D-Type", 2499.99))
	AddLicense(NewLicense("C-Type", 5999.99))
	AddLicense(NewLicense("B-Type", 19999.99))
	AddLicense(NewLicense("A-Type", 25999.99))
	AddLicense(NewLicense("S-Type", 59999.99))
	AddLicense(NewLicense("Pro", 124999.99))
	AddLicense(NewLicense("Vintage", 74999.99))
	AddLicense(NewLicense("Ultimate", 249999.99))
}

func CreateFormulaUltimateCars() {
	formUlt := NewCar("Formula Ultimate Gen. 2", 4, 1399999.99, "FormulaUltimate2")

	AddCar(formUlt)

	formUlt.AddLicense(GetLicense("A-Type"))
}

func CreateCars() {
	starterKart := NewCar("125cc Kart", 9, 1005.99, "125cc Kart")
	shiftKart := NewCar("125cc Shifter Kart", 8, 1925.99, "125cc Shifter Kart")
	superKart := NewCar("250cc Super Kart", 10, 5948.99, "250cc Super Kart")
	fVee := NewCar("Formula Vee", 99, 19362.99, "Formula Vee")
	mclarenlm := NewCar("McLaren F1 LM", 35, 89319.99, "Street Car")

	AddCar(shiftKart)
	AddCar(starterKart)
	AddCar(superKart)
	AddCar(fVee)
	AddCar(mclarenlm)

	rookieLic := GetLicense("Rookie")
	dLic := GetLicense("D-Type")

	starterKart.AddLicense(rookieLic)
	shiftKart.AddLicense(rookieLic)
	fVee.AddLicense(rookieLic)
	superKart.AddLicense(dLic)
}

func SetupLicenses() {
	//Setup Cars & Laptimes
	rookieLic := GetLicense("Rookie")
	dLic := GetLicense("D-Type")

	rookieLic.car = *cars[3]
	dLic.car = *cars[4]
}

func SetNationalities() {
	nationalities := [...]string{"United States of America", "Canada", "Germany"}
	NATIONS = nationalities[:]
	fmt.Println(NATIONS)
}

func SetupSeasonLengths() {
	SEASON_LENGTHS = append(SEASON_LENGTHS, 5)
	SEASON_LENGTHS = append(SEASON_LENGTHS, 7)
	SEASON_LENGTHS = append(SEASON_LENGTHS, 9)
	SEASON_LENGTHS = append(SEASON_LENGTHS, 12)
	SEASON_LENGTHS = append(SEASON_LENGTHS, 15)
	SEASON_LENGTHS = append(SEASON_LENGTHS, 22)
}

func SetupRaceLengths() {
	RACE_LENGTHS = append(RACE_LENGTHS, 5)
	RACE_LENGTHS = append(RACE_LENGTHS, 13)
	RACE_LENGTHS = append(RACE_LENGTHS, 25)
	RACE_LENGTHS = append(RACE_LENGTHS, 48)
	RACE_LENGTHS = append(RACE_LENGTHS, 66)
}

func SetupDefaultSettings() {
	HARDCORE = false
	DIFFICULTY = 90.0
	TIRE_WEAR = 3
	CAR_DAMAGE = 1
	SEASON_LENGTH = 2
	RACE_LENGTH = 2
}

var SETUP_FUNCS []func() = []func(){SetupSeasonLengths, SetupRaceLengths, SetupDefaultSettings, SetNationalities, CreateTracks, CreateLicenses, CreateCars, CreateFormulaUltimateCars, SetupLicenses}

func RunSetup() {
	for _, fun := range SETUP_FUNCS {
		fun()
	}
}

func SetupTest() {
	//Settings
	SetupSeasonLengths()
	SetupRaceLengths()
	SetupDefaultSettings()

	//Gameplay
	SetNationalities()
	CreateTracks()
	CreateLicenses()
	CreateCars()
	SetupLicenses()

	fmt.Println(cars[0])
}
