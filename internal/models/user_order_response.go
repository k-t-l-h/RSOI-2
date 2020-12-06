package models

import (
	"github.com/google/uuid"
	"time"
)

type UserOrderResponse struct {
	OrderUUID      uuid.UUID `json:"orderUid"`
	Date           time.Time `json:"date"`
	Model          string    `json:"model"`
	Size           string    `json:"size"`
	WarrantyDate   time.Time `json:"warrantyDate"`
	WarrantyStatus string    `json:"warrantyStatus"`
}

// User Order Response
// swagger:response UserOrdersResponse
type SwaggerUserOrdersResponse struct {
	// in:body
	Body []UserOrderResponse
}

// User Order Response
// swagger:response UserOrderResponse
type SwaggerUserOrderResponse struct {
	// in:body
	Body UserOrderResponse
}