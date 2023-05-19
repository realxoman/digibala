package models

type AdminRank struct {
	Id    int
	Title string
}

type Administrator struct {
	Id      int
	User_id int
	Rank    AdminRank
	User    User
}
