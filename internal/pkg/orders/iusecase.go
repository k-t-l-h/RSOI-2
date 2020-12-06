package orders

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IUseCase interface {
	GetOneOrder(uuid.UUID, uuid.UUID) (models.Orders, uint)
	GetOrder(uuid.UUID) (models.Orders, uint)
	GetOrders(uuid.UUID) ([]models.Orders, uint)
	MakeOrder(uuid.UUID, models.OrderResponse) uint
}
