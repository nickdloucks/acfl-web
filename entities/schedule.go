package entities

type MatchName string

const (
	// --- Pre-Season: ---
	DT1      MatchName = "DT Round 1"
	DT2      MatchName = "DT Round 2"
	DT3      MatchName = "DT Round 3"
	ACFTR1   MatchName = "ACFT Round 1"
	ACFTQF   MatchName = "ACFT Quarterfinal"
	ACFTSemi MatchName = "ACFT Semi-Final"
	ACFTCG   MatchName = "ACFT Championship"
	// --- Regular Season: ---
	Game1  MatchName = "Game 1"
	Game2  MatchName = "Game 2"
	Game3  MatchName = "Game 3"
	Game4  MatchName = "Game 4"
	Game5  MatchName = "Game 5"
	Game6  MatchName = "Game 6"
	Game7  MatchName = "Game 7"
	Game8  MatchName = "Game 8"
	Game9  MatchName = "Game 9"
	Game10 MatchName = "Game 10"
	// --- Post-Season: ---
	PineBowl   MatchName = "Pine Bowl"
	AnchorBowl MatchName = "Anchor Bowl"
	AxBowl     MatchName = "Ax Bowl"
	AlpineBowl MatchName = "Alpine Bowl"
)

var StringToMatchName = map[string]MatchName{
	// --- Pre-Season: ---
	"DT Round 1":        DT1,
	"DT Round 2":        DT2,
	"DT Round 3":        DT3,
	"ACFT Round 1":      ACFTR1,
	"ACFT Quarterfinal": ACFTQF,
	"ACFT Semi-Final":   ACFTSemi,
	"ACFT Championship": ACFTCG,
	// --- Regular Season: ---
	"Game 1":  Game1,
	"Game 2":  Game2,
	"Game 3":  Game3,
	"Game 4":  Game4,
	"Game 5":  Game5,
	"Game 6":  Game6,
	"Game 7":  Game7,
	"Game 8":  Game8,
	"Game 9":  Game9,
	"Game 10": Game10,
	// --- Post-Season: ---
	"Pine Bowl":   PineBowl,
	"Anchor Bowl": AnchorBowl,
	"Ax Bowl":     AxBowl,
	"Alpine Bowl": AlpineBowl,
}

// TO-DO: a Shedule is really just an ordered list of MatchResults
// it could MAYBE be implemented as a DB table where each game title is a column
// and the vlaue is an id ref to a MatchResult/Event in that table...

type Schedule struct {
	Id   string `json:"id"`   // primary key, uuid
	Team string `json:"team"` // foreign key, subject team for this instance of schedule
	Year uint16 `json:"year"`
	// TO-DO: should these lists of MatchResult objects be foreign keys as well to the match results "fact" table?
	// should they be expanded as named references to specific potential match slots to preserve relational structure?
	// PreSeasonOpponents  []MatchResult   `json:"pre_season_opponents"`
	// RegSeasonOpponents  [10]MatchResult `json:"reg_season_opponents"`
	// PostSeasonOpponents []MatchResult   `json:"post_season_opponents"`
}

// TO-DO: maybe this is a table where each column is a MatchName and each row is a TeamYear...?

func (s *Schedule) SchedulePreSeason() {
	// TO-DO: instantiates event objects for preseason matches
}

func (s *Schedule) ScheduleRegSeason() {
	// TO-DO: instantiates event objects for regular season matches
}

func (s *Schedule) SchedulePostSeason() {
	// TO-DO: instantiates event objects for prostseason matches
}
