package entities

// PLAYER SINGLE-GAME STAT-LINE:

type PlayerGameStatLine struct {
	// --- Ids:
	Id           string // primary-key
	GameEventId  string // (foreign key)
	PlayerYearId string // (foreign key)
	TeamYearId   string // (foreign key)

	// --- Single-Game CSV stats file output:
	// Player info:
	Position  PlayerPosition
	FirstName string
	LastName  string
	// ---------Stats:-------------
	// NOTE: any yardage stat could be positive or negative.
	// Most non-yardage quantities are >= 0
	// Stats with "Longest" in the name refer to yardage on single plays,
	// thus the following contraints apply:  -120 <= LongestX <= 120
	// In the most extreme case, a single play could only extend from the back of one endzone to the back of the other
	QBCompletions       uint8
	QBAttempts          uint8
	QBPassYards         int16
	QBPassTDs           uint8
	QBInts              uint8
	QBLongestPass       int8
	QBTimesSacked       uint8
	RushAttempts        uint8
	RushYards           int16
	RushTDs             uint8
	LongestRush         int8
	Fumbles             uint8
	Receptions          uint8
	ReceivingYards      int16
	ReceivingTDs        uint8
	LongestReception    int8
	YardsAfterCatch     int16
	Drops               uint8
	Targets             uint8
	Tackles             uint8
	Sacks               uint8
	Interceptions       uint8
	KnockDowns          uint8
	ForcedFumbles       uint8
	FumbleRecoveries    uint8
	FumbleRecoveryYards int16
	DefensiveTDs        uint8
	TacklesForLoss      uint8
	Safeties            uint8
	Penalties           uint8
	PenaltyYards        int16
	FGMade              uint8
	FGAttempted         uint8
	XPMade              uint8
	XPAttempted         uint8
	Punts               uint8
	PuntsInside20       uint8
	PuntTouchbacks      uint8
	PuntYards           int16
	PuntReturns         uint8
	PuntReturnYards     int16
	PuntReturnTDs       uint8
	KickReturnYards     int16
	KickReturns         uint8
	KickReturnTDs       uint8
}
