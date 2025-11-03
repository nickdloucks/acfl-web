package entities

import(
	"testing"
)

func TestValidateMatchName(t *testing.T) {
	t.Run("returns error if input is bad",  func(t *testing.T) {
		_, got := validateMatchName("bad input")
		if got == nil {
			t.Errorf("got nil, want error")
		}
	})
	t.Run("returns formatted string and nil if input is good",  func(t *testing.T) {
		got, err := validateMatchName("ax bowl")
		if err != nil {
			t.Errorf("got %s from error return val, want nil", err.Error())
		}
		want := "Ax Bowl"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}