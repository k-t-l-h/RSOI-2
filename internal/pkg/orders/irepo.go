package orders

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IRepo interface {
	SaveOrder(uuid.UUID, models.OrderResponse) uint
	ReadOrder(uuid.UUID, uuid.UUID) (models.Orders, uint)
	ReadItem(uuid.UUID) (models.Orders, uint)
	ReadOrders(uuid.UUID) ([]models.Orders, uint)
}
