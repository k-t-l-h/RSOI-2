package models

import uuid "github.com/google/uuid"

// swagger:model
type OrderItems struct {
	ID            int       `json:"id"`
	Cancelled     bool      `json:"cancelled"`
	OrderItemUuid uuid.UUID `json:"order_item_uuid"`
	OrderUuid     uuid.UUID `json:"order_uuid"`
	ItemID        int       `json:"item_id"`
}
