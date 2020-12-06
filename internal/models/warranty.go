package models

import (
	uuid "github.com/google/uuid"
	"time"
)


type Warranty struct {
	ID       int
	Comment  string
	ItemUuid uuid.UUID `json:"itemUid"`
	Status   string    `json:"status"`
	Warranty time.Time `json:"warrantyDate"`
}

// Warranty
// swagger:model Warranty
type SwaggerWarranty struct {
	// in:body
	Body Warranty
}