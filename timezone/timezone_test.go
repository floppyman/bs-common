package timezone

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("UTC accepted", func(t *testing.T) {
		if err := Validate("UTC"); err != nil {
			t.Errorf("UTC should be valid, got: %v", err)
		}
	})

	t.Run("Europe/Copenhagen accepted", func(t *testing.T) {
		if err := Validate("Europe/Copenhagen"); err != nil {
			t.Errorf("Europe/Copenhagen should be valid, got: %v", err)
		}
	})

	t.Run("empty string rejected", func(t *testing.T) {
		err := Validate("")
		if err == nil {
			t.Error("empty timezone should fail validation")
		}
	})

	t.Run("invalid zone rejected", func(t *testing.T) {
		err := Validate("Not/AZone")
		if err == nil {
			t.Error("invalid zone should fail validation")
		}
	})

	t.Run("too long rejected", func(t *testing.T) {
		tz := strings.Repeat("a", 251)
		err := Validate(tz)
		if err == nil {
			t.Error("over-250 runes should fail validation")
		}
	})
}
