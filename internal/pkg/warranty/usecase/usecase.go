package repo

import (
	"github.com/google/uuid"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/warranty"
)

type WrntUsecase struct {
	repo warranty.IRepo
}

func NewWrntUsecase(repo warranty.IRepo) *WrntUsecase {
	return &WrntUsecase{repo: repo}
}

func (r *WrntUsecase) GetInfo(uuid uuid.UUID) (models.Warranty, uint) {
	w, code := r.repo.ReadWarranty(uuid)
	return w, code
}

func (r *WrntUsecase) GetWarrantyResult(uuid uuid.UUID,
	req models.WarrantyRequest) (models.WarrantyResponse, uint) {

	w, code := r.repo.ReadWarranty(uuid)
	if code != models.OKAY {
		return models.WarrantyResponse{}, code
	}

	resp := models.WarrantyResponse{Date: w.Warranty.String()}
	if w.Status != "ON_WARRANTY" {
		resp.Decision = "REFUSED"
		return resp, models.OKAY
	} else {
		if req.Available == 0 {
			resp.Decision = "RETURN"
		} else {
			resp.Decision = "FIXING"
		}
	}

	return resp, code
}

func (r *WrntUsecase) StartWarranty(uuid uuid.UUID) uint {
	return r.repo.CreateWarranty(uuid)

}

func (r *WrntUsecase) DeleteWarranty(uuid uuid.UUID) uint {
	return r.repo.DeleteWarranty(uuid)
}
