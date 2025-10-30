package entities

type PlayerYear struct {
	Id   string `json:"id"` // primary key
	Year uint16 `json:"year"`
	Team string `json:"team"` // foreign key
}
