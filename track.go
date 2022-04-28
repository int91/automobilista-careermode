package main

type Track struct {
	name     string
	length   float64
	turns    int
	rentCost float64
	dlc      bool
}

func NewTrack(name string, length float64, turns int, rentCost float64, dlc bool) *Track {
	t := Track{name: name, length: length, turns: turns, rentCost: rentCost, dlc: dlc}
	return &t
}

func AddTrack(t *Track) {
	tracks = append(tracks, t)
}
