package main

import "fmt"

type License struct {
	name     string
	testCost float64
	car      *Car
	track    *Track
	laptime  float64
}

func NewLicense(name string, testCost float64) *License {
	newLic := License{name: name, testCost: testCost}
	return &newLic
}

func AddLicense(lic *License) {
	licenses = append(licenses, lic)
}

func GetLicense(name string) *License {
	for _, i := range licenses {
		if i.name == name {
			return i
		}
	}
	return nil
}

func (lic *License) DrawLicenseTestSetupScreen() {
	//TODO: Draw "Press 1 once loaded into the time trial"
	//TODO: If the player does not own the car for the test, they may rent or purchase it.
	fmt.Println(fmt.Sprintf("Please load Practice on %s with the %s", lic.track.name, lic.car.name))
	fmt.Println(fmt.Sprintf("You will be given 3 laps to beat a laptime of %f", lic.laptime))
}
