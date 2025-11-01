package entities

// import "github.com/nickdloucks/tournaments"

// An event
type MatchEvent struct {
	Id          string    `json:"id"`        // primary key
	AwayRank    uint8     `json:"away_rank"` // away team's rank/seed at beginning of match
	AwayScore   uint8     `json:"away_score"`
	AwayTeam    string    `json:"away_team"` // foreign key: Team.name
	HomeRank    uint8     `json:"home_rank"` // home team's rank/seed at beginning of match
	HomeScore   uint8     `json:"home_score"`
	HomeTeam    string    `json:"home_team"`    // foreign key: Team.name
	MatchName   MatchName `json:"match_name"`   // indicates the match's place in the schedule for the year
	NeutralSite bool      `json:"neutral_site"` // indicates whether the game was played at a neutral site
	// if the match was played at a neutral site, then the "Home" team is the team with the better rank entering the match
	Overtime uint8          `json:"overtime"` // default is 0; counts ovetime periods if there are 1 or more
	Status   CompUnitStatus `json:"status"`
	Year     uint16         `json:"year"`
}
