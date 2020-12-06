package store

import (
	"github.com/google/uuid"
)

type IRepo interface {
	CheckUser(uuid uuid.UUID) uint
}
