package entities

import(
	"testing"
)

func TestValidatePlayerPosition(t *testing.T) {
	t.Run("return error if invalid player position", func(t *testing.T) {
		got := validatePlayerPosition("invalid string")
		if got == nil {
			t.Errorf("got nil, want an error")
		}
	})
	t.Run("returns nil if string is a valid position", func(t *testing.T){
		got := validatePlayerPosition("QB")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}