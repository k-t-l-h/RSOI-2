package warehouse

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IUseCase interface {
	GetItemInfo(uuid uuid.UUID) (models.Items, uint)
	TakeItem(items models.OrderRequest) (models.OrderResponse, uint)
	DeleteItem(uuid uuid.UUID) uint
}
