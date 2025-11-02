package usecases

import(
	"nickdloucks/acfl-web/internal/entities"
)

type RosterRepository interface {
	Create() error
	FindById(id entities.UuidV7Str) (entities.Roster, error)
	FindByTeamYear(ty entities.TeamYear) (entities.Roster, error)
	AddPlayerYearById(pyId entities.UuidV7Str) error
	RemovePlayerYearById(pyId entities.UuidV7Str) error
	BulkAddPlayerYear(py []entities.PlayerYear) error
	BulkRemovePlayerYear(py []entities.PlayerYear) error
	Finalize() error
}
