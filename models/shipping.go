package models

type Shipping struct {
	ID        int
	ProductID int
	AddressID int
	Timestamp string
	Type      string
}
