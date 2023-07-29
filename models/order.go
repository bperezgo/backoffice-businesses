package models

import "time"

type Order struct {
	ID                    string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	UserId                string
	DeliveryInformationId string
}

type OrderProduct struct {
	OrderId string
	Sku     string
}
