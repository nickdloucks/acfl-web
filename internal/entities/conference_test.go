package entities

import (
	"testing"
)

func TestValidateConfName(t *testing.T) {
	t.Run("returns error with bad input name", func(t *testing.T){
		got := validateConfName("fake conference")
		if got == nil {
			t.Errorf("got %s, want an error", got)
		}
	})
	t.Run("returns no error when given a valid name", func(t *testing.T) {
		got := validateConfName("Marine Conference")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}

func TestValidateConfAbrv(t *testing.T) {
	t.Run("returns error with bad input abbreviation", func(t *testing.T){
		got := validateConfAbrv("fake conference abrv")
		if got == nil {
			t.Errorf("got %s, want an error", got)
		}
	})
	t.Run("returns no error when given a valid abbreviation", func(t *testing.T) {
		got := validateConfAbrv("MC")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}

func TestValidateConfTitleGameStr(t *testing.T) {
	t.Run("returns error with bad input title game", func(t *testing.T){
		got := validateConfTitleGameStr("fake conference title game")
		if got == nil {
			t.Errorf("got %s, want an error", got)
		}
	})
	t.Run("returns no error when given a valid title game", func(t *testing.T) {
		got := validateConfTitleGameStr("Ax Bowl")
		if got != nil {
			t.Errorf("got %s, want nil", got.Error())
		}
	})
}