package entities

type Roster struct { // not a database table: it's a filtered list of players
	// Id      string   `json:"id"`   // primary key, uuid
	Team    string   `json:"team"` // subject team, foreign key, unique
	Players []Player `json:"players"`
}

// maybe a Roster is really just a "Team" filter applied to a list of PlayerYear objects
// and therefore does not need an Id or a database table