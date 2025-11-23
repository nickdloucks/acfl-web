package entities

import (
	"errors"
	"strings"
)

type Conference struct {
	Id               UuidV7Str `json:"id"`   // primary key, uuid
	Name             string    `json:"name"` // unique
	Abrv             string    `json:"abrv"` // unique, abbreviation for the conference name
	ChampionshipGame MatchName `json:"championship_game"`
	ConfAdjustableConfig
}

type ConfAdjustableConfig struct {
	Color1 ColorStr `json:"color_1"` // hex color code
	Color2 ColorStr `json:"color_2"` // hex color code
	Color3 ColorStr `json:"color_3"` // hex color code
}

// --- Conference factory func with validation ---
func NewConference(
	id UuidV7Str,
	name string,
	abrv string,
	ccg string,
	config ConfAdjustableConfig,
) (*Conference, error) {
	abrv = strings.ToUpper(abrv)
	if id == "" {
		return nil, errors.New("conference id cannot be empty")
	}
	newConfName, confNameErr := capitalizeWords(name)
	if confNameErr != nil {
		return nil, confNameErr
	}

	return &Conference{
		Id:               id,
		Name:             newConfName,
		Abrv:             strings.ToUpper(abrv),
		ChampionshipGame: MatchName(ccg),
		ConfAdjustableConfig: ConfAdjustableConfig{
			Color1: validateColor(string(config.Color1)),
			Color2: validateColor(string(config.Color2)),
			Color3: validateColor(string(config.Color3)),
		},
	}, nil
}
