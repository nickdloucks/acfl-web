package entities


type SingleElimBracket interface {
	GetParticipants() []MatchParticipant
	GetMatches() []MatchEvent

}

// Expects a list of participants that is pre-sorted ascending from "best" to "worst" seed
func ScheduleBracket(participants []MatchParticipant) SingleElimBracket {
	NewSingleElimBracket()
}
