package warranty

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IRepo interface {
	CreateWarranty(uuid.UUID) uint
	ReadWarranty(uuid.UUID) (models.Warranty, uint)
	DeleteWarranty(uuid.UUID) uint
}
