package usecases

import(
	"nickdloucks/acfl-web/internal/entities"
)


type TeamRepository interface {
	Create(name string) error
	BulkCreate(names []string) error
	FindById(id entities.UuidV7Str) (entities.Team, error)
}


func NewTeam(p entities.UuidV7Provider, name string) *entities.Team {
	return &entities.Team{
		Id: p.NewUuidV7(),
		Name: name,
	}
}