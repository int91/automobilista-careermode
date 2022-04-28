package main

import (
	"fmt"
)

var NATIONS []string

var F_ULT_ENTRY_FEE float64 = 899999.99

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

func CreateRookieCars() {
	starterKart := NewCar("125cc Kart", 9, 1005.99, "125cc Kart")
	shiftKart := NewCar("125cc Shifter Kart", 8, 1925.99, "125cc Shifter Kart")
	superKart := NewCar("250cc Super Kart", 10, 5948.99, "250cc Super Kart")
	AddCar(shiftKart)
	AddCar(starterKart)
	AddCar(superKart)
	starterKart.AddLicense(GetLicense("Rookie"))
	shiftKart.AddLicense(GetLicense("Rookie"))
	superKart.AddLicense(GetLicense("D-Type"))
}

func SetNationalities() {
	nationalities := [...]string{"United States of America", "Canada", "Germany"}
	NATIONS = nationalities[:]
	fmt.Println(NATIONS)
}

func SetupTest() {
	SetNationalities()
	CreateLicenses()
	CreateRookieCars()
	fmt.Println(cars[0])
}
