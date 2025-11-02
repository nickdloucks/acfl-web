package usecases

import(
	"nickdloucks/acfl-web/internal/entities"
)

func NewTeam(p entities.UuidV7Provider, name string) *entities.Team {
	return &entities.Team{
		Id: p.NewUuidV7(),
		Name: name,
	}
}