package models

import uuid "github.com/google/uuid"

// swagger:model
type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	UserUuid uuid.UUID `json:"user_uuid"`
}
