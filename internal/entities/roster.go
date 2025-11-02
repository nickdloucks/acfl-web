package entities

type Roster struct { // not a database table: it's a filtered list of players
	Id          UuidV7Str    `json:"id"`
	TeamYearId  UuidV7Str    `json:"team_year_id"` // subject team, foreign key, unique
	PlayerYears []PlayerYear `json:"players"`
	Finalized   bool         `json:"finalized"`
}

type RosterRepository interface {
	Create() error
	FindById(id UuidV7Str) (Roster, error)
	FindByTeamYear(ty TeamYear) (Roster, error)
	AddPlayerYearById(pyId UuidV7Str) error
	RemovePlayerYearById(pyId UuidV7Str) error
	BulkAddPlayerYear(py []PlayerYear) error
	BulkRemovePlayerYear(py []PlayerYear) error
	Finalize() error
}
