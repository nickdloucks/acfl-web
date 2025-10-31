package entities

type Team struct {
	// note: the Team object transcends "Year".
	// Player stats, team stats, and game results are assigned to a year
	Id         string `json:"id"`   // primary key, uuid
	Abrv       string `json:"abrv"`       // unique, abbreviation for the team name
	Name       string `json:"name"` // unique
}
