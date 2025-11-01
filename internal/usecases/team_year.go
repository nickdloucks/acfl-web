package usecases

import (
	"nickdloucks/acfl-web/internal/entities"
)

func NewInitialTeamYear(team entities.Team, year uint16, config entities.TeamYearUpdatableConfig) entities.TeamYear {
	return entities.TeamYear{
		Id: UuidV7.NewUuidv7(),
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