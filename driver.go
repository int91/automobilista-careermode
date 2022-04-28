package main

type Driver struct {
	name      string
	number    int
	file_name string
	licenses  []*License
}

func (d *Driver) AssignNewNumber(num int) {
	d.number = num
}
