package entities

type Conference struct {
	Id               string    `json:"id"`   // primary key, uuid
	Name             string    `json:"name"` // unique
	Abrv             string    `json:"abrv"` // unique, abbreviation for the conference name
	ChampionshipGame MatchName `json:"championship_game"`
	Color1           string    `json:"color_1"`
	Color2           string    `json:"color_2"`
	Color3           string    `json:"color_3"`
}
