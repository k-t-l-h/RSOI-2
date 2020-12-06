package repo

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/pkg/store"
)

type StrUsecase struct {
	repo store.IRepo
}

func NewStrUsecase(repo store.IRepo) *StrUsecase {
	return &StrUsecase{repo: repo}
}

func (r *StrUsecase) Check(uuid uuid.UUID) uint {
	return r.repo.CheckUser(uuid)
}
