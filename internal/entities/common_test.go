package entities

import (
	"testing"
)

func TestValidateColor(t *testing.T) {
	t.Run("returns default color if string too long", func(t *testing.T) {
		got := validateColor("1234567890")
		want := ColorStr("#363636")
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("returns default color if input contains non-hex chars", func(t *testing.T) {
		got := validateColor("#36363g")
		want := ColorStr("#363636")
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("returns formatted input if valid", func(t *testing.T) {
		got := validateColor(string(WHITE))
		want := ColorStr("#ffffff")
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
