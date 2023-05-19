package models
import "time"

type Brand struct {
    ID                 int
    Name               string
    Description        string
    Logo               string
    CountryID          int
	CreatedAt          time.Time
    UpdatedAt          time.Time
}
