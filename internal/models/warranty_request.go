package models

type WarrantyRequest struct {
	Reason    string `json:"reason"`
	Available int    `json:"available,omitempty"`
}

// WarrantyRequest
// swagger:model WarrantyRequest
type SwaggerWarrantyRequest struct {
	// in:body
	Body WarrantyRequest
}
