package main

type Team struct {
	name string
	cars [2]*Car
}

type RaceTeam struct {
	name  string
	teams []*Team
}

func NewTeam(name string) *Team {
	nTeam := Team{name: name}
	return &nTeam
}
