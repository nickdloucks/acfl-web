package entities

type Player struct {
	// note: the Player object does NOT transcend "Year".
	// Player stats, team stats, and game results are assigned to a year
	Id    string `json:"id"`    // primary key, uuid
	First string `json:"first"` // first name
	Last  string `json:"last"`
	// TO-DO: fields for skills?
}
