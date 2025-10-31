package entities

type Team struct {
	// note: the Team object transcends "Year".
	// Player stats, team stats, and game results are assigned to a year
	Id   string `json:"id"`   // primary key, uuid
	Name string `json:"name"` // unique
}
