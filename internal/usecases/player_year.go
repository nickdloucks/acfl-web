package usecases

import "nickdloucks/acfl-web/internal/entities"


type PlayerYearRepository interface {
	CreateInitialPlayerYear(player entities.Player) (entities.PlayerYear, error)
	CreateSubsequentPlayerYear(player entities.Player) (entities.PlayerYear, error)
	AssignToRoster(py entities.PlayerYear) error
}