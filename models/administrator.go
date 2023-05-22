package models

type AdminRank struct {
	Id    int
	Title string
}

type Administrator struct {
	Id     int
	UserId int
	Rank   AdminRank
	User   User
}
