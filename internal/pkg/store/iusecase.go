package store

import (
	"github.com/google/uuid"
)

type IUseCase interface {
	Check(uuid uuid.UUID) uint
}
