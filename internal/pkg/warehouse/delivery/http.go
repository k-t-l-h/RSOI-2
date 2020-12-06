package delivery

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"lab2-microservices-k-t-l-h/internal/pkg/warehouse"
	"net/http"
)

type WrhsHandler struct {
	WarrantyAddr string
	uc           warehouse.IUseCase
}

func NewWrhsHandler(warrantyAddr string, uc warehouse.IUseCase) *WrhsHandler {
	return &WrhsHandler{WarrantyAddr: warrantyAddr, uc: uc}
}


// swagger:operation GET /api/v1/warehouse/{UUID} warehouse GetWhInfo
// ---
// summary: All orders info
//
// description: Just get all info
//
// parameters:
// - name: UUID
//   in: path
//   description: order item UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/definitions/OrderRequest"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
func (h *WrhsHandler) GetItemInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	it, code := h.uc.GetItemInfo(id)
	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, it)
	case models.NOTFOUND:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("item %s not found", ids),
		}
		utils.Response(w, http.StatusNotFound, answer)
	default:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("item %s was incorrect", ids),
		}
		utils.Response(w, http.StatusBadRequest, answer)
	}

}

// swagger:operation DELETE /api/v1/warehouse/{UUID}} warehouse DelWhInfo
// ---
// summary: All orders info
//
// description: Just get all info
//
// parameters:
// - name: UUID
//   in: path
//   description: order item UUID
//   type: string
//   required: true
// responses:
//   "204":
//     description: "Returned"
func (h *WrhsHandler) ReturnItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	code := h.uc.DeleteItem(id)
	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusNoContent, nil)
	default:
		utils.Response(w, http.StatusBadRequest, nil)
	}
}



// swagger:operation POST /api/v1/warehouse/ warehouse getWarehouseItem
// ---
// summary: Get Warranty
//
// description: Just take warranty response
//
// parameters:
// - name: Request
//   in: body
//   description: Order Request
//   schema:
//     "$ref": "#/definitions/OrderRequest"
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/OrderResponse"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
func (h *WrhsHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	//получение модели из body
	item := &models.OrderRequest{}
	easyjson.UnmarshalFromReader(r.Body, item)
	it, code := h.uc.TakeItem(*item)

	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, it)
	case models.NOTFOUND:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("order %s not found", item.OrderUUID),
		}
		utils.Response(w, http.StatusNotFound, answer)
	case models.UNAVAILABLE:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("order %s not avaliable", item.OrderUUID),
		}
		utils.Response(w, http.StatusConflict, answer)
	default:
		answer := models.ErrorMessage{
			Message: fmt.Sprintf("order %s was incorrect", item.OrderUUID),
		}
		utils.Response(w, http.StatusBadRequest, answer)
	}

}



// swagger:operation POST /api/v1/warehouse/{itemUUID}/warranty warehouse getWarehouseWarranty
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
//   description: Order Request
//   schema:
//     "$ref": "#/definitions/WarrantyRequest"
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/WarrantyOderResponse"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
func (h *WrhsHandler) GetItemWarranty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	//обращение к сервису

	/// api/v1/warranty/{itemUid}/warranty

	rs := models.WarrantyRequest{}
	easyjson.UnmarshalFromReader(r.Body, &rs)

	bd, _ := easyjson.Marshal(rs)

	addr := h.WarrantyAddr + fmt.Sprintf("/api/v1/warranty/%s/warranty", ids)
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(bd))
	if err != nil || resp == nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		utils.Response(w, http.StatusOK, nil)
	case http.StatusNotFound:
		answer := &models.ErrorMessage{
			fmt.Sprintf("Warranty not found for itemUid '%s'", ids),
		}
		easyjson.UnmarshalFromReader(r.Body, answer)
		utils.Response(w, http.StatusNotFound, answer)
	case http.StatusBadRequest:
		answer := &models.ErrorMessage{
			fmt.Sprintf("Warranty incorrect for %s", ids),
		}
		easyjson.UnmarshalFromReader(r.Body, answer)
		utils.Response(w, http.StatusBadRequest, answer)
	default:
		utils.Response(w, http.StatusUnprocessableEntity, nil)

	}

}
