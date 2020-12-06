package models

type ErrorMessage struct {
	Message string `json:"message"`
}

// Error Message
// swagger:response ErrorResponse
type SwaggerError struct {
	// in:body
	Body ErrorMessage
}
