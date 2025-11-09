package entities

type TimeRemainingStr string // format 00:00

type GameQuarter string

const (
	Q1  GameQuarter = "Q1"
	Q2  GameQuarter = "Q2"
	Q3  GameQuarter = "Q3"
	Q4  GameQuarter = "Q4"
	OT1 GameQuarter = "OT1"
	OT2 GameQuarter = "OT2"
	OT3 GameQuarter = "OT3"
	OT4 GameQuarter = "OT4"
)

type PlayEndReason string

const (
	Tackle     PlayEndReason = "Tackle"
	Incomplete PlayEndReason = "Incomplete"
	OOB        PlayEndReason = "OOB"
	OOBEndzone PlayEndReason = "OOB_In_Endzone"
	TD         PlayEndReason = "Touchdown"
	FG         PlayEndReason = "MadeFG" // find way to differentiate between FG and XP
	// missed FG a possible outcome
	// 2-pt conv made: may need to differentiate from Touchdown
)

// When processing events from the game logs, do an initial loop thru each play in order
// and create a new uuidv7 for it before putting it in a queue: this way they are all tagged in order.
// Then, a goroutine can pull off of that queue and do some more advanced analysis of each play to
// generate the rest of the PlayEvent fields and the order of plays will still be preserved, or at least
// able to be easily re-sorted later if necessary.

type PlayBeginState struct {
	StartTime             TimeRemainingStr `json:"start_time"`
	TeamInPossessionStart string           `json:"team_in_possession_start"`
	StartYdLine           int8             `json:"start_yd_line"`
	CurrentDown           uint8            `json:"current_down"`
	CurrentDistance       uint8            `json:"current_distance"`
	Offense               string           `json:"offense"` // Team.Abrv
	Defense               string           `json:"defense"` // Team.Abrv
	OffensePlaySelection  string           `json:"off_play_selection"`
	DefensePlaySelection  string           `json:"def_play_selection"`
}

type PlayEndState struct {
	EndTime             TimeRemainingStr `json:"end_time"`
	TeamInPossessionEnd string           `json:"team_in_possession_end"`
	EndYdLine           int8             `json:"end_yd_line"`
	NextDown            uint8            `json:"next_down"`
	NextDistance        uint8            `json:"next_distance"`
	PlayEndReason       PlayEndReason    `json:"play_end_reason"`
}

type PlayEvent struct {
	Id                 UuidV7Str   `json:"id"`      // primary key
	GameId             UuidV7Str   `json:"game_id"` // foreign key
	Home               string      `json:"home"`    // Team.Abrv foreign key
	Away               string      `json:"away"`    // Team.Abrv foreign key
	Quarter            GameQuarter `json:"quarter"`
	IsSpecialTeamsPlay bool        `json:"is_special_teams_play"` // default false
	Turnovers          uint8       `json:"turnovers"`             // values of 0 or 1 can be parsed as a bool, while higher numbers can still be used in the rare scenario where there are multiple turnovers on the same play
	PointsScored       uint8       `json:"points_scored"`
	PlayBeginState
	PlayEndState
	RushPlay   bool `json:"rush_play"`
	PassPlay   bool `json:"pass_play"`
	PlayAction bool `json:"play_action"`
	BlitzPlay  bool `json:"blitz_play"`
	NetYards   int8 `json:"netYards"` // since this is a single play event, this number should fit in int8, but may need conversion for aggragating into game stats
}
