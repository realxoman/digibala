package models

import "time"

type GiftCard struct {
	ID             int
	Title          string
	ShopID         string
	CurrencyID     int
	OriginalAmount int
	ExpirationDate time.Time
	Description    string
}
