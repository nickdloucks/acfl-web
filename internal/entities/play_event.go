package entities

type TimeRemainingStr string //

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

// When processing events from the game logs, do an initial loop thru each play in order
// and create a new uuidv7 for it before putting it in a queue: this way they are all tagged in order.
// Then, a goroutine can pull off of that queue and do some more advanced analysis of each play to
// generate the rest of the PlayEvent fields and the order of plays will still be preserved, or at least
// able to be easily re-sorted later if necessary.

type PlayBeginState struct {
	StartTime        TimeRemainingStr `json:"start_time"`
	TeamInPossession string           `json:"team_in_possession_strt"`
	StartingYdLine   int8             `json:"starting_yd_line"`
	Down             uint8            `json:"start_down"`
	Distance         uint8            `json:"start_distance"`
}

type PlayEndState struct {
	EndTime          TimeRemainingStr `json:"end_time"`
	TeamInPossession string           `json:"team_in_possession_end"`
	EndingYdLine     int8             `json:"ending_yd_line"`
	Down             uint8            `json:"end_down"`
	Distance         uint8            `json:"end_distance"`
}

type PlayEvent struct {
	Id        UuidV7Str   `json:"id"`
	Quarter   GameQuarter `json:"quarter"`
	Turnovers uint8       `json:"turnovers"` // values of 0 or 1 can be parsed as a bool, while higher numbers can still be used in the rare scenario where there are multiple turnovers on the same play
	PlayBeginState
	PlayEndState
}
