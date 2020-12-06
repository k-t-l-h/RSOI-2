package models

import "github.com/google/uuid"

type OrderRequest struct {
	OrderUUID uuid.UUID `json:"orderUid,omitempty"`
	Model     string    `json:"model"`
	Size      string    `json:"size"`
}

// Order Request
//swagger:model OrderRequest
type SwaggerOrderRequest struct {
	// This text will appear as description of your request body.
	// in:body
	Body OrderRequest
}
