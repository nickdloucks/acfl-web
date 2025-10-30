package entities

type Team struct {
	// note: the Team object transcends "Year".
	// Player stats, team stats, and game results are assigned to a year
	Id         string `json:"id"`   // primary key, uuid
	Name       string `json:"name"` // unique
	City       string `json:"home_city"`
	Abrv       string `json:"abrv"`       // unique, abbreviation for the team name
	Conference string `json:"conference"` // foreign key
	Color1     string `json:"color_1"`    // hex color code
	Color2     string `json:"color_2"`    // hex color code
	Color3     string `json:"color_3"`    // hex color code
	Color4     string `json:"color_4"`    // hex color code
	// Logo1Path  string `json:"logo_1_path"` 
	// Logo2Path  string `json:"logo_2_path"`
	// probably don't need these ^ paths, can just have the ui construct the path using some sort of naming convention
}
