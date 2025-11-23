package entities

type Player struct {
	// Player stats, team stats, and game results are assigned to a year
	Id    UuidV7Str `json:"id"`    // primary key, uuid
	First string    `json:"first"` // first name
	Last  string    `json:"last"`
}


