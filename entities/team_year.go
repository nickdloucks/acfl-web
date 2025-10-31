package entities

type TeamYear struct {
	Id         string `json:"id"` // primary key
	Year       uint16 `json:"year"`
	Name       string `json:"name"` // unique, format: YYYY-Team Name
	City       string `json:"home_city"`
	Conference string `json:"conference"` // foreign key
	Color1     string `json:"color_1"`    // hex color code
	Color2     string `json:"color_2"`    // hex color code
	Color3     string `json:"color_3"`    // hex color code
	Color4     string `json:"color_4"`    // hex color code
	// Roster   string `json:"roster"`   // list of PlayerYear objects
	// Schedule string `json:"schedule"` // list of MatchResult objects
	// Stats    string `json:"stats"`    // list of TeamMatchStats objects
	// note: team stats for the whole year should be dynamically calulated so that they are auto-updating
}
