package models

import "time"

type Product struct {
	Sku            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Description    string
	Price          float32
	Stock          int
	AvailableMedia AvailableMedia
}

type AvailableMedia struct {
	Images []string
	Videos []string
}
