package models

import (
	"github.com/google/uuid"
	"time"
)

// swagger:model
type OrderResponse struct {
	OrderUUID     uuid.UUID `json:"orderUid"`
	OrderItemUUID uuid.UUID `json:"orderItemUid"`
	OrderDate     time.Time
	Model         string `json:"model"`
	Size          string `json:"size"`
}

//easyjson:json
type AllOrders []OrderResponse

// Order Response
// swagger:response OrderResponse
type SwaggerOrderResponse struct {
	// in:body
	Body OrderResponse
}

// All Orders
// swagger:response AllOrders
type SwaggerAllOrders struct {
	// in:body
	Body AllOrders
}
