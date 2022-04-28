package main

type License struct {
	name     string
	testCost float64
	car      Car
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
