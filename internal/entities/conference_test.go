package entities

import (
	"testing"
)

func TestvalidateConfName(t *testing.T) {
	t.Run("returns a valid name", func(t *testing.T) {

		got := validateConfName("Marine Conference")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}

// func TestNewConference(t *testing.T) {
// 	got := NewConference(
// 		UuidV7Str("uuid-v7-string-here"),
// 		"test-name",
// 		"Ax Bowl",
// 	)
// }
