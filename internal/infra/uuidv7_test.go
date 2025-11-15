package infra

import(
	"testing"
)

func TestNewUuidV7(t *testing.T) {
	t.Run("should not be empty", func(t *testing.T) {
		if NewUuidV7() == "" {
			t.Errorf("got empty string, want uuid V7 string")
		}
	})
	t.Run("should be unique", func(t *testing.T) {
		first := NewUuidV7()
		second := NewUuidV7()
		if first == second {
			t.Errorf("two uuid V7 strings should not match. Got %s and %s", first, second)
		}
	})
}