package repo

import (
	"context"
	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"lab2-microservices-k-t-l-h/internal/models"
)

type StrRepo struct {
	pool pgxpool.Pool
}

func NewStrRepo(pool pgxpool.Pool) *StrRepo {
	return &StrRepo{pool: pool}
}

func (r *StrRepo) CheckUser(uuid uuid.UUID) uint {
	GetUser := "SELECT id FROM users WHERE user_uid = $1;"
	UserID := 0

	row := r.pool.QueryRow(
		context.Background(),
		GetUser,
		uuid,
	)

	err := row.Scan(&UserID)
	if err != nil {
		return models.NOTFOUND
	}

	return models.OKAY
}
