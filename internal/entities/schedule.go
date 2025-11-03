package entities

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

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

func validateMatchName(name string) (validatedStr string, err error) {
	if name == "" {
		return "", errors.New("invalid match name: empty")
	}
	parts := strings.Split(name, " ")
	capitalizedParts := []string{}
	for _, item := range parts {
		r, size := utf8.DecodeRuneInString(item)
		if r == utf8.RuneError {
			return "", errors.New("invalid match name")
		}
		item = string(unicode.ToUpper(r)) + item[size:]
		capitalizedParts = append(capitalizedParts, item)
	}
	name = strings.Join(capitalizedParts, " ")
	if _, ok := StringToMatchName[name]; ok {
		return name, nil
	}
	return "", errors.New("invalid match name")
}

// TO-DO: a Shedule is really just an ordered list of MatchResults
// it could MAYBE be implemented as a DB table where each game title is a column
// and the vlaue is an id ref to a MatchResult/Event in that table...

type Schedule struct {
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
	// get list of all 16 teams sorted by end-of-year rankings from previous year
	// instantiate a 16-team bracket with that list
	// - make all the MatchEvent objects for the tournament
	// - 
}

func (s *Schedule) ScheduleRegSeason() {
	// TO-DO: instantiates event objects for regular season matches
	// for each conference:
	// 		get a list of all 8 TeamYear objects
	// 		schedule a round-robin tournament with all the TeamYear objects in the conference
	// 			^ note TO-DO: there must be some mechanism for determining home/away teams w/o using a ranking
						// - lookup the relevant conf matchup between the teams from prev year, and switch the home/away
						// - if no conf matchup prev year, then "flip a coin"

	// -------------
	// figure out how to schedule non-conference games and still include a few cross-conference rivalries every year
}

func (s *Schedule) SchedulePostSeason() {
	// TO-DO: instantiates event objects for prostseason matches
	// only once top two teams in each conference have clinched:
	// create a single-elim bracket tournament:
	// ensure that, regardless of actual league ranking, when feeding in ranked list to the bracket constructor,
	// one conference must put its teams in the 1 & 4 seed slots, while the other puts them in the  2 & 3 slots
	// so that the correct teams play each other in the "semi-final" round
	// also need to override the league champ game in that tournament to be neutral site.
	// --------then:
	// only after two teams have clinched the next highest spots:
	// schedule a single MatchEvent for the Pine Bowl
}
