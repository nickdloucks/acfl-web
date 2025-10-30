package entities

type TeamYear struct {
	Id   string `json:"id"` // primary key
	Year uint16 `json:"year"`
	Name string `json:"name"` // unique, format: YYYY-Team Name
	// Roster   string `json:"roster"`   // list of PlayerYear objects
	// Schedule string `json:"schedule"` // list of MatchResult objects
	// Stats    string `json:"stats"`    // list of TeamMatchStats objects
	// note: team stats for the whole year should be dynamically calulated so that they are auto-updating
}
