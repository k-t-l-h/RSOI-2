package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"lab2-microservices-k-t-l-h/internal/models"
)

type WrhsRepo struct {
	pool pgxpool.Pool
}

func NewWrhsRepo(pool pgxpool.Pool) *WrhsRepo {
	return &WrhsRepo{pool: pool}
}

func (r *WrhsRepo) SelectItem(uuid uuid.UUID) (models.Items, uint) {
	it := models.Items{}

	SelItem := "SELECT order_item.item_id " +
		"FROM order_item " +
		"WHERE order_item.order_item_uid  = $1;"

	row := r.pool.QueryRow(context.Background(), SelItem, uuid)

	err := row.Scan(&it.ID)
	if err != nil {
		return models.Items{}, models.NOTFOUND
	}
	SelItemInfo := "SELECT model, size " +
		"FROM items " +
		"WHERE id= $1;"

	row = r.pool.QueryRow(context.Background(), SelItemInfo, it.ID)

	err = row.Scan(&it.Model, &it.Size)
	if err != nil {
		return models.Items{}, models.NOTFOUND
	}

	return it, models.OKAY
}

func (r *WrhsRepo) ReserveItem(items models.OrderRequest) (models.OrderResponse, uint) {
	//сделать запрос на информацию
	it := models.Items{}
	CheckNum := "SELECT id, available_count " +
		"FROM items " +
		"WHERE model = $1 AND size = $2;"

	queryRow := r.pool.QueryRow(context.Background(),
		CheckNum, items.Model, items.Size)

	if err := queryRow.Scan(&it.ID, &it.Available); err != nil {
		return models.OrderResponse{}, models.BADREQUEST
	}

	if it.Available <= 0 {
		return models.OrderResponse{}, models.UNAVAILABLE
	}

	//если достаточно, то начать транзакцию
	tx, err := r.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return models.OrderResponse{}, models.BADREQUEST
	}
	TakeItem := "UPDATE items " +
		"SET available_count=available_count-1 " +
		"WHERE model = $1 AND size = $2;"

	tag, err := tx.Exec(context.Background(), TakeItem, items.Model, items.Size)
	if err != nil || tag.RowsAffected() == 0 {
		_ = tx.Rollback(context.Background())
		return models.OrderResponse{}, models.BADREQUEST
	}

	MakeOrder := "INSERT INTO order_item " +
		"(canceled, order_item_uid, order_uid, item_id) " +
		"VALUES ($1, $2, $3, $4) RETURNING order_item_uid;"

	queryRow = tx.QueryRow(context.Background(),
		MakeOrder, false, uuid.New(), items.OrderUUID, it.ID)

	order := models.OrderResponse{
		OrderUUID: items.OrderUUID,
		Model:     items.Model,
		Size:      items.Size,
	}

	err = queryRow.Scan(&order.OrderItemUUID)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return models.OrderResponse{}, models.BADREQUEST
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.OrderResponse{}, models.BADREQUEST
	}

	return order, models.OKAY
}

func (r *WrhsRepo) ReturnItem(uuid uuid.UUID) uint {

	tx, err := r.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return models.BADREQUEST
	}

	CancelOrder := "UPDATE order_item " +
		"SET canceled=true WHERE order_item_uid = $1 " +
		"RETURNING item_id;"
	row := tx.QueryRow(context.Background(), CancelOrder, uuid)

	id := 0
	err = row.Scan(&id)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return models.BADREQUEST
	}

	ReturnItem := "UPDATE  items " +
		"SET available_count=available_count+1 " +
		"WHERE id = $1;"

	tag, err := tx.Exec(context.Background(), ReturnItem, id)
	if err != nil || tag.RowsAffected() == 0 {
		_ = tx.Rollback(context.Background())
		return models.BADREQUEST
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.BADREQUEST
	}

	return models.OKAY
}
