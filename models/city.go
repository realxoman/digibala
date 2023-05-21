package models

type state struct {
	ID        int
	Name      string
	CountryId int
}

type province struct {
	ID      int
	Name    string
	StateId int
}

type City struct {
	ID         int
	Name       string
	ProvinceId int
}
