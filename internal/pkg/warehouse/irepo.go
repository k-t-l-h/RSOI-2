package warehouse

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IRepo interface {
	SelectItem(uuid uuid.UUID) (models.Items, uint)
	ReserveItem(items models.OrderRequest) (models.OrderResponse, uint)
	ReturnItem(uuid uuid.UUID) uint
}
