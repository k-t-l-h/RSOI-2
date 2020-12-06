package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"lab2-microservices-k-t-l-h/internal/models"
)

type OdrRepo struct {
	pool pgxpool.Pool
}

func NewOdrRepo(pool pgxpool.Pool) *OdrRepo {
	return &OdrRepo{pool: pool}
}

func (r *OdrRepo) SaveOrder(id uuid.UUID, order models.OrderResponse) uint {
	MakeOrder := "INSERT INTO orders(" +
		" item_uid, order_date, order_uid, status, user_uid) " +
		"VALUES ( $1, $2, $3, $4, $5) RETURNING id;"

	row := r.pool.QueryRow(context.Background(), MakeOrder,
		order.OrderItemUUID, order.OrderDate, order.OrderUUID, "PAID", id)

	var sid int
	err := row.Scan(&sid)
	if err != nil {
		return models.BADREQUEST
	}

	return models.OKAY
}

func (r *OdrRepo) ReadOrder(id uuid.UUID, oid uuid.UUID) (models.Orders, uint) {
	Order := "SELECT item_uid, order_date, order_uid, status " +
		"FROM orders WHERE user_uid = $1 AND order_uid = $2;"

	row := r.pool.QueryRow(context.Background(), Order, id, oid)
	it := models.Orders{}

	it.UserUuid = id
	it.OrderUuid = oid
	err := row.Scan(&it.ItemUuid, &it.OrderDate, &it.OrderUuid, &it.Status)
	if err != nil {
		return models.Orders{}, models.NOTFOUND
	}

	return it, models.OKAY
}

func (r *OdrRepo) ReadOrders(id uuid.UUID) ([]models.Orders, uint) {
	its := []models.Orders{}

	Order := "SELECT item_uid, order_date, order_uid, status " +
		"FROM orders WHERE user_uid = $1;"

	row, err := r.pool.Query(context.Background(), Order, id)
	if err != nil {
		return []models.Orders{}, models.BADREQUEST
	}

	for row.Next() {
		it := models.Orders{}
		it.UserUuid = id
		err := row.Scan(&it.ItemUuid, &it.OrderDate, &it.OrderUuid, &it.Status)
		if err != nil {
			return []models.Orders{}, models.NOTFOUND
		}
		its = append(its, it)
	}

	return its, models.OKAY
}

func (r *OdrRepo) ReadItem(oid uuid.UUID) (models.Orders, uint) {
	Order := "SELECT item_uid  " +
		"FROM orders WHERE order_uid = $1;"

	row := r.pool.QueryRow(context.Background(), Order, oid)
	it := models.Orders{}

	err := row.Scan(&it.ItemUuid)
	if err != nil {
		return models.Orders{}, models.NOTFOUND
	}

	return it, models.OKAY
}
