package models

type Location struct {
	Lat float64
	Lng float64
}

type Address struct {
	ID       int
	UserID   int
	Title    string
	Location Location
	Address  string
	No       string
	Floor    string
}
