package delivery

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/orders"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"net/http"
)

type OdrHandler struct {
	uc     orders.IUseCase
	whAddr string
	wrAddr string
}

func NewOdrHandler(uc orders.IUseCase, whAddr string, wrAddr string) *OdrHandler {
	return &OdrHandler{uc: uc, whAddr: whAddr, wrAddr: wrAddr}
}

// swagger:operation POST /api/v1/orders/{UUID} order MakeOrder
// ---
// summary: Make Order
//
// description: Just Make One Order
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
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
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
func (h *OdrHandler) MakeOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	orderID := uuid.New()
	order := &models.OrderRequest{}
	easyjson.UnmarshalFromReader(r.Body, order)
	order.OrderUUID = orderID

	body, _ := easyjson.Marshal(order)

	addr := h.whAddr + fmt.Sprintf("/api/v1/warehouse/")
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))

	if err != nil || resp == nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		order := models.OrderResponse{}
		easyjson.UnmarshalFromReader(resp.Body, &order)
		h.uc.MakeOrder(id, order)
		addr := h.wrAddr + fmt.Sprintf("/api/v1/warranty/%s", order.OrderItemUUID)
		http.Post(addr, "application/json", nil)
		utils.Response(w, http.StatusOK, order)
	default:
		utils.CopyResponse(w, resp)
	}

}

// swagger:operation GET /api/v1/orders/{UUID}/{orderUUID} order GetOrder
// ---
// summary: One order info
//
// description: Just get one order info
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
// - name: orderUUID
//   in: path
//   description: user order UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/Orders"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
func (h *OdrHandler) OrderInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	oids, _ := vars["orderUUID"]
	oid, _ := uuid.Parse(oids)

	//получение информации о заказе
	order, code := h.uc.GetOneOrder(id, oid)
	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, order)
	case models.NOTFOUND:
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("order %s not found", ids)}
		utils.Response(w, http.StatusNotFound, msg)
	default:
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("order %s was incorrect", oids)}
		utils.Response(w, http.StatusBadRequest, msg)
	}
}

// swagger:operation GET /api/v1/orders/{UUID} order GetOrders
// ---
// summary: All orders info
//
// description: Just get all info
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/AllOrders"
func (h *OdrHandler) UsersOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	order, code := h.uc.GetOrders(id)
	switch code {
	case models.OKAY:
		utils.Response(w, http.StatusOK, order)
	case models.NOTFOUND:
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("orders for user %s not found", ids)}
		utils.Response(w, http.StatusNotFound, msg)
	default:
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("orders %s are incorrect", ids)}
		utils.Response(w, http.StatusBadRequest, msg)
	}

}

// swagger:operation Delete /api/v1/orders/{UUID} order returnOrder
// ---
// summary: Return Order
//
// description: Return Order
//
// parameters:
// - name: UUID
//   in: path
//   description: orders UUID
//   type: string
//   required: true
// responses:
//   "204":
//	   description: Success
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
func (h *OdrHandler) ReturnOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	id, _ := uuid.Parse(ids)

	order, code := h.uc.GetOrder(id)
	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("orders for user %s not found", ids)}
		utils.Response(w, http.StatusNotFound, msg)
	}

	addr := h.whAddr + fmt.Sprintf("/api/v1/warehouse/%s", order.ItemUuid)

	req, err := http.NewRequest(http.MethodDelete, addr, bytes.NewBufferString(""))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.Body == nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		utils.Response(w, http.StatusNoContent, nil)
	case http.StatusBadRequest:
		utils.Response(w, http.StatusBadRequest, nil)
	default:
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("this was unexpected, error code %d",
				resp.StatusCode)}
		utils.Response(w, http.StatusUnprocessableEntity, msg)
	}

}

// swagger:operation POST /api/v1/orders/{orderUUID}/warranty order getWarranty
// ---
// summary: Get Warranty
//
// description: Just take warranty response
//
// parameters:
// - name: orderUUID
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
func (h *OdrHandler) GetWarranty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["orderUUID"]
	id, _ := uuid.Parse(ids)

	rs := models.WarrantyRequest{}
	easyjson.UnmarshalFromReader(r.Body, &rs)

	bd, _ := easyjson.Marshal(rs)

	order, code := h.uc.GetOrder(id)
	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("orders for user %s not found", ids)}
		utils.Response(w, http.StatusNotFound, msg)
	}

	addr := h.wrAddr + fmt.Sprintf("/api/v1/warranty/%s/warranty",
		order.ItemUuid.String())

	resp, err := http.Post(addr, "application/json",
		bytes.NewBuffer(bd))

	if err != nil || resp == nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		des := models.WarrantyResponse{}
		easyjson.UnmarshalFromReader(resp.Body, &des)
		desd := models.WarrantyOderResponse{
			Date:     des.Date,
			Decision: des.Decision,
		}
		utils.Response(w, http.StatusOK, desd)
	default:
		utils.CopyResponse(w, resp)
	}

}
