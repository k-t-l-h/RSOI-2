package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"lab2-microservices-k-t-l-h/internal/models"
	"time"
)

type WrntRepo struct {
	pool pgxpool.Pool
}

func NewWrntRepo(pool pgxpool.Pool) *WrntRepo {
	return &WrntRepo{pool: pool}
}

func (r *WrntRepo) CreateWarranty(id uuid.UUID) uint {

	InsertWarranty := "INSERT INTO warranty( " +
		"comment, item_uid, status, warranty_date) " +
		"VALUES ($1, $2, $3, $4) RETURNING id;"

	w := models.Warranty{
		ItemUuid: id,
		Warranty: time.Now(),
		Status:   "ON_WARRANTY",
		Comment:  "NEW THING",
	}

	row := r.pool.QueryRow(context.Background(), InsertWarranty, w.Comment,
		w.ItemUuid, w.Status, w.Warranty)

	err := row.Scan(&w.ID)
	if err != nil {
		return models.BADREQUEST
	}

	return models.OKAY
}

func (r *WrntRepo) ReadWarranty(id uuid.UUID) (models.Warranty, uint) {
	GetWarranty := "SELECT id, " +
		"comment, " +
		"item_uid, " +
		"status, " +
		"warranty_date FROM warranty " +
		"WHERE item_uid = $1;"

	w := models.Warranty{
		ItemUuid: id,
	}

	row := r.pool.QueryRow(context.Background(), GetWarranty, w.ItemUuid)
	err := row.Scan(&w.ID, &w.Comment, &w.ItemUuid, &w.Status, &w.Warranty)
	if err != nil {
		return w, models.NOTFOUND
	}
	return w, models.OKAY
}

func (r *WrntRepo) DeleteWarranty(id uuid.UUID) uint {

	DeleteWarranty := "UPDATE warranty " +
		"SET status=$1 WHERE item_uid = $2;"

	w := models.Warranty{
		ItemUuid: id,
		Status:   "REMOVED_FROM_WARRANTY",
	}

	exec, err := r.pool.Exec(context.Background(), DeleteWarranty, w.Status, w.ItemUuid)

	if exec.RowsAffected() == 0 {
		return models.NOTFOUND
	}

	if err != nil {
		return models.BADREQUEST
	}

	return models.OKAY
}
