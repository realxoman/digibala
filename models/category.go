package models

type Category struct {
	Id       int32
	Title    string
	Products []Product
}
