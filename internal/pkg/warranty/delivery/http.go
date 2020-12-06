package delivery

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"lab2-microservices-k-t-l-h/internal/pkg/warranty"
	"net/http"
)

type WrntHandler struct {
	uc warranty.IUseCase
}

func NewWrntHandler(uc warranty.IUseCase) *WrntHandler {
	return &WrntHandler{uc: uc}
}

// swagger:operation GET /api/v1/warranty/{itemUUID} warranty GetWrInfo
// ---
// summary: Check warranty status
//
// description: Check warranty status
//
// parameters:
// - name: itemUUID
//   in: path
//   description: item UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/definitions/Warranty"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
func (h *WrntHandler) Info(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, err := uuid.Parse(ids)

	if err != nil {
		utils.Response(w, http.StatusBadRequest, nil)
	}

	resp, code := h.uc.GetInfo(id)

	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, resp)
	case models.NOTFOUND:
		answer := models.ErrorMessage{Message: fmt.Sprintf("warranty with id %s not found", ids)}
		utils.Response(w, http.StatusNotFound, answer)
	default:
		answer := models.ErrorMessage{Message: fmt.Sprintf("bad request for warranty with id %s", ids)}
		utils.Response(w, http.StatusBadRequest, answer)

	}
}


// swagger:operation POST /api/v1/warranty/{itemUUID}/warranty warranty getWarranty
// ---
// summary: Get Warranty
//
// description: Just take warranty response
//
// parameters:
// - name: itemUUID
//   in: path
//   description: order UUID
//   type: string
//   required: true
// - name: Request
//   in: body
//   description: Warranty Request
//   schema:
//     "$ref": "#/definitions/WarrantyRequest"
//   required: true
// responses:
//   "200":
//     "$ref": "#/definitions/Warranty"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
func (h *WrntHandler) InfoResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	req := &models.WarrantyRequest{}
	easyjson.UnmarshalFromReader(r.Body, req)

	resp, code := h.uc.GetWarrantyResult(id, *req)

	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, resp)
	case models.NOTFOUND:
		answer := models.ErrorMessage{Message: fmt.Sprintf("warranty with id %s not found", ids)}
		utils.Response(w, http.StatusNotFound, answer)
	default:
		answer := models.ErrorMessage{Message: fmt.Sprintf("bad request for warranty with id %s", ids)}
		utils.Response(w, http.StatusBadRequest, answer)
	}
}


// swagger:operation POST /api/v1/warranty/{itemUUID} warranty startWarranty
// ---
// summary: Start Warranty
//
// description: Start warranty
//
// parameters:
// - name: itemUUID
//   in: path
//   description: order UUID
//   type: string
//   required: true
// responses:
//   "204":
//     description: "Started"
func (h *WrntHandler) StartWarranty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	code := h.uc.StartWarranty(id)

	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusNoContent, nil)
	default:
		answer := models.ErrorMessage{Message: fmt.Sprintf("bad request for warranty with id %s", ids)}
		utils.Response(w, http.StatusBadRequest, answer)
	}

}

// swagger:operation DELETE /api/v1/warranty/{itemUUID} warranty endWarranty
// ---
// summary: End Warranty
//
// description: End warranty
//
// parameters:
// - name: itemUUID
//   in: path
//   description: order UUID
//   type: string
//   required: true
// responses:
//   "204":
//     description: "Ended"
func (h *WrntHandler) EndWarranty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	code := h.uc.DeleteWarranty(id)

	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusNoContent, nil)
	default:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("bad request for warranty with id %s",
				ids)}
		utils.Response(w, http.StatusBadRequest, answer)
	}
}
