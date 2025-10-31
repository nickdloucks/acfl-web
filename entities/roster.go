package entities

type Roster struct { // not a database table: it's a filtered list of players
	TeamYear    TeamYear     `json:"team_year"` // subject team, foreign key, unique
	PlayerYears []PlayerYear `json:"players"`
}