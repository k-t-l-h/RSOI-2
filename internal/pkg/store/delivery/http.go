package delivery

import (
	"bytes"
	"fmt"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	_ "github.com/swaggo/http-swagger"
	"lab2-microservices-k-t-l-h/internal/models"
	"lab2-microservices-k-t-l-h/internal/pkg/store"
	"lab2-microservices-k-t-l-h/internal/pkg/utils"
	"net/http"
)

type StrHandler struct {
	orderAddr string
	whAddr    string
	wrAddr    string
	uc        store.IUseCase
}

func NewStrHandler(orderAddr string,
	whAddr string,
	wrAddr string,
	uc store.IUseCase) *StrHandler {
	return &StrHandler{orderAddr: orderAddr,
		whAddr: whAddr,
		wrAddr: wrAddr,
		uc:     uc}
}

func (h StrHandler) CheckUser(uuids string) uint {
	id, err := uuid.Parse(uuids)
	if err != nil {
		return models.BADREQUEST
	}
	code := h.uc.Check(id)
	return code
}


// swagger:operation GET /api/v1/store/{UUID}/orders store getOrders
// ---
// summary: Get All Orders
//
// description: Get all orders info
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/UserOrdersResponse"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
//   "500":
//     "$ref": "#/responses/ErrorResponse"
func (h *StrHandler) Orders(w http.ResponseWriter, r *http.Request) {

	//проверка корректности профиля
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	code := h.CheckUser(ids)

	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("no user with id %s", ids),
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	//запрос на другой микросервис
	addr := h.orderAddr + fmt.Sprintf("/api/v1/orders/%s", ids)
	resp, err := http.Get(addr)
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}

	//обработка ответа
	switch resp.StatusCode {
	case http.StatusInternalServerError:
		msg := models.ErrorMessage{Message: "service unavailable"}
		utils.Response(w, http.StatusUnprocessableEntity, msg)
	default:
		utils.CopyResponse(w, resp)
		resp.Body.Close()
	}
}

// swagger:operation GET /api/v1/store/{UUID}/{orderUUID} store getOrder
// ---
// summary: Get One Order
//
// description: Get one order info
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
// - name: orderUUID
//   in: path
//   description: order UUID
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/UserOrderResponse"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
//   "500":
//     "$ref": "#/responses/ErrorResponse"
func (h *StrHandler) OrdersInfo(w http.ResponseWriter, r *http.Request) {

	//проверка пользователя
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	code := h.CheckUser(ids)

	answer := models.UserOrderResponse{}

	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("no user with id %s", ids),
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	//получение информации о заказе
	orderIds, _ := vars["ORDER_UUID"]

	//запрос на статус и дату
	addr := h.orderAddr + fmt.Sprintf("/api/v1/orders/%s/%s", ids, orderIds)
	resp, err := http.Get(addr)

	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode != http.StatusOK {
		utils.CopyResponse(w, resp)
		resp.Body.Close()
		return
	}

	orders := &models.Orders{}
	easyjson.UnmarshalFromReader(resp.Body, orders)
	resp.Body.Close()

	id, _ := uuid.Parse(orderIds)
	answer.OrderUUID = id
	answer.Date = orders.OrderDate

	//запрос на данные по itemIds
	addr = h.whAddr + fmt.Sprintf("/api/v1/warehouse/%s", orders.ItemUuid)
	resp, err = http.Get(addr)
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode != http.StatusOK {
		utils.CopyResponse(w, resp)
		resp.Body.Close()
		return
	}

	item := &models.Items{}
	easyjson.UnmarshalFromReader(resp.Body, item)
	resp.Body.Close()

	answer.Model = item.Model
	answer.Size = item.Size

	//запрос на гарантию
	addr = h.wrAddr + fmt.Sprintf("/api/v1/warranty/%s", orders.ItemUuid)
	resp, err = http.Get(addr)
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	if resp.StatusCode != http.StatusOK {
		utils.CopyResponse(w, resp)
		resp.Body.Close()
		return
	}
	warr := &models.Warranty{}
	easyjson.UnmarshalFromReader(resp.Body, warr)
	resp.Body.Close()

	answer.WarrantyDate = warr.Warranty
	answer.WarrantyStatus = warr.Status

	utils.Response(w, http.StatusOK, answer)
}


// swagger:operation POST /api/v1/store/{UUID}/{orderUUID}/warranty store getOrderWarranty
// ---
// summary: Get One Order
//
// description: Get one order warranty
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
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
//     "$ref": "#/responses/UserOrderResponse"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
//   "500":
//     "$ref": "#/responses/ErrorResponse"
func (h *StrHandler) Warranty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	code := h.CheckUser(ids)

	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("no user with id %s", ids),
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	orderIds, _ := vars["ORDER_UUID"]
	orderID, _ := uuid.Parse(orderIds)

	wr := &models.WarrantyRequest{}
	easyjson.UnmarshalFromReader(r.Body, wr)
	body, _ := easyjson.Marshal(wr)

	addr := fmt.Sprintf(h.orderAddr+"/api/v1/orders/%s/warranty", orderIds)
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		answer := &models.WarrantyOderResponse{}
		easyjson.UnmarshalFromReader(resp.Body, answer)
		answer.UUID = orderID
		utils.Response(w, http.StatusOK, answer)
	default:
		utils.CopyResponse(w, resp)
	}

}


// swagger:operation POST /api/v1/store/{UUID}/purchase store purchaseOrder
// ---
// summary: Get One Order
//
// description: Get one order warranty
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
//   "201":
//     description: "Success"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
//   "500":
//     "$ref": "#/responses/ErrorResponse"
func (h *StrHandler) Purchase(w http.ResponseWriter, r *http.Request) {

	//проверка пользователя
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	code := h.CheckUser(ids)

	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("no user with id %s", ids),
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	//проверка запроса
	item := &models.Items{}
	err := easyjson.UnmarshalFromReader(r.Body, item)
	if err != nil {
		msg := models.ErrorMessage{
			Message: "incorrect json",
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	//запрос на микросервис
	body, _ := easyjson.Marshal(item)
	addr := h.orderAddr + fmt.Sprintf("/api/v1/orders/%s", ids)
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	defer resp.Body.Close()
	//проверка ответа на запрос
	switch resp.StatusCode {
	case http.StatusOK:
		it := models.OrderResponse{}
		easyjson.UnmarshalFromReader(resp.Body, &it)
		w.Header().Set("Location",
			fmt.Sprintf("/%s", it.OrderUUID.String()))

		utils.Response(w, http.StatusCreated, nil)
	default:
		utils.CopyResponse(w, resp)
	}

}


// swagger:operation DELETE /api/v1/store/{UUID}/{orderUUID}/refund store refundOrder
// ---
// summary: Get One Order
//
// description: Get one order warranty
//
// parameters:
// - name: UUID
//   in: path
//   description: user UUID
//   type: string
//   required: true
// - name: orderUUID
//   in: path
//   description: order UUID
//   type: string
//   required: true
// responses:
//   "204":
//     description: "Returned"
//   "400":
//     "$ref": "#/responses/ErrorResponse"
//   "404":
//     "$ref": "#/responses/ErrorResponse"
//   "409":
//     "$ref": "#/responses/ErrorResponse"
//   "422":
//     "$ref": "#/responses/ErrorResponse"
//   "500":
//     "$ref": "#/responses/ErrorResponse"
func (h *StrHandler) Refund(w http.ResponseWriter, r *http.Request) {

	//проверка пользователя
	vars := mux.Vars(r)
	ids, _ := vars["UUID"]
	orderIds, _ := vars["ORDER_UUID"]
	code := h.CheckUser(ids)

	if code != models.OKAY {
		msg := models.ErrorMessage{
			Message: fmt.Sprintf("no user with id %s", ids),
		}
		utils.Response(w, http.StatusNotFound, msg)
		return
	}

	//запрос на микросервис
	addr := h.orderAddr + fmt.Sprintf("/api/v1/orders/%s", orderIds)
	req, err := http.NewRequest(http.MethodDelete, addr, bytes.NewBufferString(""))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, nil)
		return
	}
	defer resp.Body.Close()
	utils.CopyResponse(w, resp)
}
