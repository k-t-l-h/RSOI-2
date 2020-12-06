package repo

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/orders"
)

type OdrUsecase struct {
	repo orders.IRepo
}

func NewOdrUsecase(repo orders.IRepo) *OdrUsecase {
	return &OdrUsecase{repo: repo}
}

func (r OdrUsecase) MakeOrder(id uuid.UUID, od models.OrderResponse) uint {
	return r.repo.SaveOrder(id, od)
}

func (r *OdrUsecase) GetOneOrder(id uuid.UUID, oid uuid.UUID) (models.Orders, uint) {
	return r.repo.ReadOrder(id, oid)
}
func (r *OdrUsecase) GetOrders(id uuid.UUID) ([]models.Orders, uint) {
	return r.repo.ReadOrders(id)
}

func (r *OdrUsecase) GetOrder(oid uuid.UUID) (models.Orders, uint) {
	return r.repo.ReadItem(oid)
}
