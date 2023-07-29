package handlers

import (
	"github.com/bperezgo/backoffice-businesses/shared/platform/handler"
)

func GetHandlers() []handler.Handler {
	getHealth := NewGetHealthHandler()
	createOrders := NewCreateOrderHandler()

	return []handler.Handler{
		getHealth,
		createOrders,
	}
}
