package infra

import(
	"testing"
)

func TestNewUuidV7(t *testing.T) {
	testUuidInstance := NewUuidV7ProviderImpl()
	t.Run("should not be empty", func(t *testing.T) {
		str, err := testUuidInstance.NewUuidV7()
		if err != nil {
			t.Errorf("got error from uuidv7 implementation, want uuidv7 string")
		}
		if str == "" {
			t.Errorf("got empty string, want uuid V7 string")
		}
	})
	t.Run("should be unique", func(t *testing.T) {
		first, _ := testUuidInstance.NewUuidV7()
		second, _ := testUuidInstance.NewUuidV7()
		if first == second {
			t.Errorf("two uuid V7 strings should not match. Got %s and %s", first, second)
		}
	})
}