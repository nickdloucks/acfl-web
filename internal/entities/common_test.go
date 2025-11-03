package entities

import (
	"testing"
)

func TestValidateColor(t *testing.T) {
	t.Run("returns default color if invalid",  func(t *testing.T){
		got := validateColor("bad input")
		want := ColorStr("#363636")
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}