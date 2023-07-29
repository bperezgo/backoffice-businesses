package models

import "time"

type User struct {
	ID               string
	Email            string
	Phone            string
	CountryPhoneCode string
	Name             string
	LastName         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
