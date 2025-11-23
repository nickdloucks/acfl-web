package usecases

import "nickdloucks/acfl-web/internal/entities"


type ConferenceRepository interface {
	FindById(id entities.UuidV7Str) (entities.Conference, error)
}