package main

type Player struct {
	driver        *Driver
	money         float64
	teamMoney     float64
	nationality   *string
	carCollection []*Car
	teams         []*Team
}

func NewPlayer(name string, number int) *Player {
	drv := Driver{name: name, number: number}
	newP := Player{driver: &drv}
	return &newP
}

func (p *Player) SetNationality(index int) {
	p.nationality = &NATIONS[index]
}

//Method *Player.Pay takes bool team and float64 amount. This will pay money based on which value team is
func (p *Player) Pay(team bool, amount float64) {
	if p.HasMoney(team, amount) {
		if team {
			p.teamMoney -= amount
		} else {
			p.money -= amount
		}
	}
}

//Method *Player.HasMoney takes bool team and float64 amount. This will return true if one account has enough money
func (p *Player) HasMoney(team bool, amount float64) bool {
	if team && p.teamMoney >= amount {
		return true
	} else if !team && p.money >= amount {
		return true
	}
	return false
}

func (p *Player) BuyTeamCar(car *Car) {
	if p.HasMoney(true, car.price) {
		p.carCollection = append(p.carCollection, car)
	}
}

func (p *Player) BuyCar(car *Car) {
	if p.HasMoney(false, car.price) {
		p.carCollection = append(p.carCollection, car)
	}
}

func (p *Player) GetDriverNumber() int {
	return p.driver.number
}

func (p *Player) AttemptDriverTest(lic *License) {
	if p.HasMoney(false, lic.testCost) {
		//TODO: Go to start game screen with various settings and such
		lic.DrawLicenseTestSetupScreen()
	}
}
