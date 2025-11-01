package usecases

import "testing"

func TestNewTeamYear(t *testing.T) {
	testTeam := Team{Name: "test_team"}
	testTY := NewTeamYear(testTeam)
	t.Run("year must not be empty", func(t *testing.T) {

	})
}