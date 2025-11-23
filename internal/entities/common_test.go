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

func TestValidateGameStatus(t *testing.T) {
	t.Run("return an error if input is invalid", func(t *testing.T) {
		got := validateGameStatus("bad input")
		if got == nil {
			t.Errorf("got nil, want an error")
		}
	})
	t.Run("return nil if input is a valid game status", func(t *testing.T) {
		got := validateGameStatus("FINAL")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}

func TestCapitalizeWords(t *testing.T) {
	got, _ := capitalizeWords("hello world")
	want := "Hello World"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}