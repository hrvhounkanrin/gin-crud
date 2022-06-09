package models

import "gorm.io/gorm"

type Insured struct {
	gorm.Model
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
}
