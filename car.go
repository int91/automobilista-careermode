package main

type Car struct {
	name     string
	number   int
	price    float64
	licenses []*License
	cType    string
	driver   *Driver
}

func NewCar(ne string, num int, price float64, cType string) *Car {
	newCar := Car{name: ne, number: num, price: price, cType: cType, licenses: nil, driver: nil}
	return &newCar
}

func AddCar(c *Car) {
	cars = append(cars, c)
}

func (car *Car) AddLicense(lic *License) {
	car.licenses = append(car.licenses, lic)
}

func (car *Car) AssignDriver(driver *Driver) {
	car.driver = driver
}

func (car *Car) GetDriver() *Driver {
	return car.driver
}

func (car *Car) DriverIsNil() bool {
	return car.driver == nil
}

func (car *Car) GetLicenses() []*License {
	return car.licenses
}

func (car *Car) DriverHasLicense(license *License) bool {
	for _, i := range car.licenses {
		if i == license {
			return true
		}
	}
	return false
}
