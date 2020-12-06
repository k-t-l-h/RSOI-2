package warranty

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
)

type IUseCase interface {
	GetInfo(uuid.UUID) (models.Warranty, uint)
	GetWarrantyResult(uuid.UUID, models.WarrantyRequest) (models.WarrantyResponse, uint)
	StartWarranty(uuid uuid.UUID) uint
	DeleteWarranty(uuid uuid.UUID) uint
}
