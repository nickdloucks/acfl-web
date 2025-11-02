package usecases

import (
	"nickdloucks/acfl-web/internal/entities"
)


type TeamYearRepository interface {
	CreateInitialYear(team entities.Team, year uint8, config entities.TeamYearUpdatableConfig) (entities.TeamYear, error)
	CreateSubsequentYear(team entities.Team, year uint8, config ...entities.TeamYearUpdatableConfig) (entities.TeamYear, error)
	FindById(id entities.UuidV7Str) (entities.TeamYear, error)
	FindByNameAndYear(name string, year uint8) (entities.TeamYear, error)
	FindByTeamIdAndYear(teamId entities.UuidV7Str, year uint8) (entities.TeamYear, error)
}

func NewInitialTeamYear(team entities.Team, year uint16, config entities.TeamYearUpdatableConfig, p UuidV7Provider) entities.TeamYear {
	return entities.TeamYear{
		Id: p.NewUuidV7(),
		TeamId: team.Id,
		Name: team.Name,
		Year: year,
		TeamYearUpdatableConfig: config,
		// Abrv: config.Abrv,
		// City: config.City,
		// Conference: config.Conference,
		// Color1: config.Color1,
		// Color2: config.Color2,
		// Color3: config.Color3,
		// Color4: config.Color4,
	}
}