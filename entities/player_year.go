package entities

type PlayerYear struct {
	Id       string `json:"id"`        // primary key
	PlayerId string `json:"player_id"` // foreign key to Player
	Year     uint16 `json:"year"`
	Team     string `json:"team"` // foreign key to TeamYear
}
