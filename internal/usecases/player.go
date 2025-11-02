package usecases

import (
	"nickdloucks/acfl-web/internal/entities"
)

type PlayerRepository interface {
	FindById(id entities.UuidV7Str) (entities.Player, error)
}