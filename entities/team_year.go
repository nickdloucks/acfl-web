package entities

type TeamYear struct {
	Id         string `json:"id"`   // primary key
	Name       string `json:"name"` // unique
	Year       uint16 `json:"year"`
	Abrv       string `json:"abrv"` // unique, abbreviation for the team name
	City       string `json:"home_city"`
	Conference string `json:"conference"` // foreign key
	Color1     string `json:"color_1"`    // hex color code
	Color2     string `json:"color_2"`    // hex color code
	Color3     string `json:"color_3"`    // hex color code
	Color4     string `json:"color_4"`    // hex color code
}
