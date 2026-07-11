package usecases

import(
	"nickdloucks/acfl-web/internal/entities"
)


type TeamRepository interface {
	// Create(name string) error
	// BulkCreate(names []string) error
	FindById(id entities.UuidV7Str) (entities.Team, error)
}


var uuidProviderUsecaseImpl = 

func NewTeam(name string) *entities.Team {
	return &entities.Team{
		Id: uuidProviderUsecaseImpl.NewUuidV7(),
		Name: name,
	}
}
// https://www.reddit.com/r/rust/comments/1try5g9/fastuuidv7_creating_uuidv7_faster_than_allocating/
// rabbit hole 6