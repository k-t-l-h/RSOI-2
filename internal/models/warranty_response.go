package models

import "github.com/google/uuid"

type WarrantyResponse struct {
	UUID     uuid.UUID `json:"-"`
	Date     string    `json:"date"`
	Decision string    `json:"decision"`
}

type WarrantyOderResponse struct {
	UUID     uuid.UUID `json:"orderUid,omitempty"`
	Date     string    `json:"warrantyDate"`
	Decision string    `json:"decision"`
}

// Warranty Oder Response
// swagger:response WarrantyOderResponse
type SwaggerWarrantyOderResponse struct {
	// in:body
	Body WarrantyOderResponse
}
