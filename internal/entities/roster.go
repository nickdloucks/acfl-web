package entities

type Roster struct { // not a database table: it's a filtered list of players
	Id          UuidV7Str    `json:"id"`
	TeamYearId  UuidV7Str    `json:"team_year_id"` // subject team, foreign key, unique
	PlayerYears []PlayerYear `json:"players"`
	Finalized   bool         `json:"finalized"`
}
