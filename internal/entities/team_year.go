package entities

// Represents a single year for a Team.
// Has a many-to-one relationship to a Team.
// Has a one-to-one relationship to a Roster.
// Has a one-to-many relationship to a PlayerYear.
type TeamYear struct {
	Id     UuidV7Str `json:"id"`      // primary key
	TeamId UuidV7Str `json:"team_id"` // foreign key
	Name   string    `json:"name"`    // unique
	Year   uint16    `json:"year"`
	TeamYearUpdatableConfig
}

type TeamYearUpdatableConfig struct {
	Abrv       string `json:"abrv"` // unique, abbreviation for the team name
	City       string `json:"home_city"`
	Conference string `json:"conference"` // foreign key
	Color1     string `json:"color_1"`    // hex color code
	Color2     string `json:"color_2"`    // hex color code
	Color3     string `json:"color_3"`    // hex color code
	Color4     string `json:"color_4"`    // hex color code
}

