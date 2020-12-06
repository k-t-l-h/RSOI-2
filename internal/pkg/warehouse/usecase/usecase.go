package repo

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/warehouse"
)

type WrhsUsecase struct {
	repo warehouse.IRepo
}

func NewWrhsUsecase(repo warehouse.IRepo) *WrhsUsecase {
	return &WrhsUsecase{repo: repo}
}

// GET /api/v1/warehouse/{orderItemUid}
func (u *WrhsUsecase) GetItemInfo(uuid uuid.UUID) (models.Items, uint) {
	return u.repo.SelectItem(uuid)
}

func (u *WrhsUsecase) TakeItem(items models.OrderRequest) (models.OrderResponse, uint) {
	return u.repo.ReserveItem(items)
}

func (u *WrhsUsecase) DeleteItem(uuid uuid.UUID) uint {
	return u.repo.ReturnItem(uuid)
}
