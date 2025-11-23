package entities

// Used to represent any team that has ever existed in the league.
// A Team has a one-to-many relationship with TeamYears
type Team struct {
	// note: the Team object transcends "Year".
	// Player stats, team stats, and game results are assigned to a year
	Id   UuidV7Str `json:"id"`   // primary key, uuid
	Name string    `json:"name"` // unique
}
