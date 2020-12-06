package models

import (
	uuid "github.com/google/uuid"
	"time"
)

type Orders struct {
	ID        int       `json:"-"`
	ItemUuid  uuid.UUID `json:"itemUid"`
	OrderDate time.Time `json:"orderDate"`
	OrderUuid uuid.UUID `json:"orderUid"`
	Status    string    `json:"status"`
	UserUuid  uuid.UUID
}

// All Orders
// swagger:response Orders
type SwaggerAOrders struct {
	// in:body
	Body Orders
}
