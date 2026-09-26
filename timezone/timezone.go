package timezone

import (
	"errors"
	"fmt"
	"time"
	_ "time/tzdata"
	"unicode/utf8"
)

const maxRunes = 250

func Validate(tz string) error {
	if tz == "" {
		return errors.New("timezone is required")
	}
	if utf8.RuneCountInString(tz) > maxRunes {
		return fmt.Errorf("timezone must be at most %d characters", maxRunes)
	}
	if _, err := time.LoadLocation(tz); err != nil {
		return fmt.Errorf("invalid timezone %q: %w", tz, err)
	}
	return nil
}
