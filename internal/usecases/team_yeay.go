package usecases

import (
	"github.com/nickdloucks/acfl-web/domain"
)

func NewInitialTeamYear(team Team, year uint16, config TeamYearUpdatableConfig) TeamYear {
	return TeamYear{
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