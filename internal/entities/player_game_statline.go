package entities

// PLAYER SINGLE-GAME STAT-LINE:

type PlayerGameStatLine struct {
	// --- Ids:
	Id           string `json:"id"` // primary-key
	GameEventId  string `json:"game_event_id"` // (foreign key)
	PlayerYearId string `json:"player_year_id"` // (foreign key)
	TeamYearId   string `json:"team_year_id"` // (foreign key)

	// --- Single-Game CSV stats file output:
	// Player info:
	Position  PlayerPosition `json:"Position"`
	FirstName string         `json:"FirstName"`
	LastName  string         `json:"LastName"`
	// ---------Stats:-------------
	// NOTE: any yardage stat could be positive or negative.
	// Most non-yardage quantities are >= 0
	// Stats with "Longest" in the name refer to yardage on single plays,
	// thus the following contraints apply:  -120 <= LongestX <= 120
	// In the most extreme case, a single play could only extend from the back of one endzone to the back of the other
	QBCompletions       uint8 `json:"QBCompletions"`
	QBAttempts          uint8 `json:"QBAttempts"`
	QBPassYards         int16 `json:"QBPassYards"`
	QBPassTDs           uint8 `json:"QBPassTDs"`
	QBInts              uint8 `json:"QBInts"`
	QBLongestPass       int8  `json:"QBLongestPass"`
	QBTimesSacked       uint8 `json:"QBTimesSacked"`
	RushAttempts        uint8 `json:"RushAttempts"`
	RushYards           int16 `json:"RushYards"`
	RushTDs             uint8 `json:"RushTDs"`
	LongestRush         int8  `json:"LongestRush"`
	Fumbles             uint8 `json:"Fumbles"`
	Receptions          uint8 `json:"Receptions"`
	ReceivingYards      int16 `json:"ReceivingYards"`
	ReceivingTDs        uint8 `json:"ReceivingTDs"`
	LongestReception    int8  `json:"LongestReception"`
	YardsAfterCatch     int16 `json:"YardsAfterCatch"`
	Drops               uint8 `json:"Drops"`
	Targets             uint8 `json:"Targets"`
	Tackles             uint8 `json:"Tackles"`
	Sacks               uint8 `json:"Sacks"`
	Interceptions       uint8 `json:"Interceptions"`
	KnockDowns          uint8 `json:"KnockDowns"`
	ForcedFumbles       uint8 `json:"ForcedFumbles"`
	FumbleRecoveries    uint8 `json:"FumbleRecoveries"`
	FumbleRecoveryYards int16 `json:"FumbleRecoveryYards"`
	DefensiveTDs        uint8 `json:"DefensiveTDs"`
	TacklesForLoss      uint8 `json:"TacklesForLoss"`
	Safeties            uint8 `json:"Safeties"`
	Penalties           uint8 `json:"Penalties"`
	PenaltyYards        int16 `json:"PenaltyYards"`
	FGMade              uint8 `json:"FGMade"`
	FGAttempted         uint8 `json:"FGAttempted"`
	XPMade              uint8 `json:"XPMade"`
	XPAttempted         uint8 `json:"XPAttempted"`
	Punts               uint8 `json:"Punts"`
	PuntsInside20       uint8 `json:"PuntsInside20"`
	PuntTouchbacks      uint8 `json:"PuntTouchbacks"`
	PuntYards           int16 `json:"PuntYards"`
	PuntReturns         uint8 `json:"PuntReturns"`
	PuntReturnYards     int16 `json:"PuntReturnYards"`
	PuntReturnTDs       uint8 `json:"PuntReturnTDs"`
	KickReturnYards     int16 `json:"KickReturnYards"`
	KickReturns         uint8 `json:"KickReturns"`
	KickReturnTDs       uint8 `json:"KickReturnTDs"`
}
