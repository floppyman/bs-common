package timeutil

import (
	"time"

	"github.com/guregu/null/v6"
)

// FromUnix returns the UTC time.Time corresponding to the given Unix epoch seconds.
// Fractional seconds are truncated toward zero.
func FromUnix(val float64) time.Time {
	return time.Unix(int64(val), 0).UTC()
}

// ToUnix returns the Unix epoch seconds for the given time.
func ToUnix(t time.Time) float64 {
	return float64(t.Unix())
}

// FromUnixPtr returns null.Time for a nil pointer, otherwise the UTC time for the seconds value.
func FromUnixPtr(val *float64) null.Time {
	if val == nil {
		return null.Time{}
	}
	return null.TimeFrom(time.Unix(int64(*val), 0).UTC())
}

// ToUnixPtr returns nil for an invalid null.Time, otherwise a pointer to its epoch seconds.
func ToUnixPtr(t null.Time) *float64 {
	if !t.Valid {
		return nil
	}
	v := float64(t.Time.Unix())
	return &v
}
