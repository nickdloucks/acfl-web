package entities

type Conference struct {
	Id       string `json:"id"`   // primary key, uuid
	Name     string `json:"name"` // unique
	Abrv     string `json:"abrv"` // unique, abbreviation for the conference name
	Color1   string `json:"color_1"`
	Color2   string `json:"color_2"`
	LogoPath string `json:"logo_path"`
}
