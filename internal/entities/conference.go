package entities

import (
	"errors"
)

type Conference struct {
	Id               UuidV7Str      `json:"id"`   // primary key, uuid
	Name             ConferenceName `json:"name"` // unique
	Abrv             ConferenceAbrv `json:"abrv"` // unique, abbreviation for the conference name
	ChampionshipGame MatchName      `json:"championship_game"`
	ConfAdjustableConfig
}

type ConfAdjustableConfig struct {
	Color1 ColorStr `json:"color_1"` // hex color code
	Color2 ColorStr `json:"color_2"` // hex color code
	Color3 ColorStr `json:"color_3"` // hex color code
}

// Enumerate allowed Conference fields that will not be editable:

// --- Conference Name ---
type ConferenceName string

const (
	MarineConfName ConferenceName = "Marine Conference"
	SpruceConfName ConferenceName = "Spruce Conference"
)

// Enumerates allowed conference names.
var ConfNameStrToConst = map[string]ConferenceName{
	"Marine Conference": MarineConfName,
	"Spruce Conference": SpruceConfName,
}

func validateConfName(name string) error {
	if _, ok := ConfNameStrToConst[name]; ok {
		return nil
	}
	return errors.New("invalid conference name string")
}

// --- Conference Abbreviation ---
type ConferenceAbrv string

const (
	MarineConfAbrv ConferenceAbrv = "MC"
	SpruceConfAbrv ConferenceAbrv = "SC"
)

// Enumerates allowed converence abbreviation strings.
var ConfAbrvStrToConst = map[string]ConferenceAbrv{
	"MC": MarineConfAbrv,
	"SC": SpruceConfAbrv,
}

func validateConfAbrv(abrv string) error {
	if _, ok := ConfAbrvStrToConst[abrv]; ok {
		return nil
	}
	return errors.New("invalid conference abbreviation string")
}

// --- Conference Title Game a.k.a Championship Game ---
type ConfTitleGame string

const (
	MCCG ConfTitleGame = "Anchor Bowl"
	SCCG ConfTitleGame = "Ax Bowl"
)

var ConfCGStrToConst = map[string]ConfTitleGame{
	"Anchor Bowl": MCCG,
	"Ax Bowl":     SCCG,
}

func validateConfTitleGameStr(ccg string) error {
	if _, ok := ConfCGStrToConst[ccg]; ok {
		return nil
	}
	return errors.New("invalid conference title game string")
}

// --- Conference factory func with validation ---
func NewConference(
	id UuidV7Str,
	name string,
	abrv string,
	ccg string,
	config ConfAdjustableConfig,
) (*Conference, error) {

	if id == "" {
		return nil, errors.New("conference id cannot be empty")
	}
	confNameErr := validateConfName(name)
	if confNameErr != nil {
		return nil, confNameErr
	}
	confAbrvErr := validateConfAbrv(abrv)
	if confAbrvErr != nil {
		return nil, confAbrvErr
	}
	confTgErr := validateConfTitleGameStr(ccg)
	if confTgErr != nil {
		return nil, confTgErr
	}

	return &Conference{
		Id:               id,
		Name:             ConfNameStrToConst[name],
		Abrv:             ConfAbrvStrToConst[abrv],
		ChampionshipGame: MatchName(ConfCGStrToConst[ccg]),
		ConfAdjustableConfig: ConfAdjustableConfig{
			Color1: validateColor(string(config.Color1)),
			Color2: validateColor(string(config.Color2)),
			Color3: validateColor(string(config.Color3)),
		},
	}, nil
}
