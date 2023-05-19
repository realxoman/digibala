package models
import "time"

type Brand struct {
    ID                 int
    BrandName          string
    Description        string
    Logo               string
    Country            int
    FounderOwner       string
    ContactInformation Contact
    Products           []string
	CreatedAt          time.Time
    UpdatedAt          time.Time
}

type Contact struct {
    ID       int
    Email    string
    Phone    string
}