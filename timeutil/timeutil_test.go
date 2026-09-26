package timeutil

import (
	"testing"
	"time"

	"github.com/guregu/null/v6"
)

func TestFromUnix(t *testing.T) {
	t.Run("zero returns Unix epoch", func(t *testing.T) {
		got := FromUnix(0)
		want := time.Unix(0, 0).UTC()
		if !got.Equal(want) {
			t.Errorf("FromUnix(0) = %v, want %v", got, want)
		}
		if got.Location() != time.UTC {
			t.Errorf("FromUnix(0) location = %v, want UTC", got.Location())
		}
	})

	t.Run("known UTC value round-trips", func(t *testing.T) {
		want := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		got := FromUnix(float64(want.Unix()))
		if !got.Equal(want) {
			t.Errorf("FromUnix = %v, want %v", got, want)
		}
	})

	t.Run("fractional seconds truncated toward zero", func(t *testing.T) {
		got := FromUnix(1735689600.9)
		if got.Unix() != 1735689600 {
			t.Errorf("fractional .9 should truncate, got %d", got.Unix())
		}
	})

	t.Run("result is in UTC regardless of system tz", func(t *testing.T) {
		got := FromUnix(1735689600)
		if got.Location() != time.UTC {
			t.Errorf("location = %v, want UTC", got.Location())
		}
	})
}

func TestToUnix(t *testing.T) {
	t.Run("now round-trips through FromUnix/ToUnix", func(t *testing.T) {
		now := time.Now().UTC().Truncate(time.Second)
		v := ToUnix(now)
		got := FromUnix(v)
		if !got.Equal(now) {
			t.Errorf("round-trip mismatch: got %v, want %v", got, now)
		}
	})
}

func TestFromUnixPtr(t *testing.T) {
	t.Run("nil returns invalid null.Time", func(t *testing.T) {
		got := FromUnixPtr(nil)
		if got.Valid {
			t.Errorf("FromUnixPtr(nil) should be invalid, got %v", got)
		}
	})

	t.Run("non-nil returns UTC time", func(t *testing.T) {
		v := 1735689600.0
		got := FromUnixPtr(&v)
		if !got.Valid {
			t.Fatal("FromUnixPtr(&v) should be valid")
		}
		if !got.Time.Equal(time.Unix(1735689600, 0).UTC()) {
			t.Errorf("time = %v, want 2025-01-01 UTC", got.Time)
		}
		if got.Time.Location() != time.UTC {
			t.Errorf("location = %v, want UTC", got.Time.Location())
		}
	})
}

func TestToUnixPtr(t *testing.T) {
	t.Run("invalid returns nil", func(t *testing.T) {
		got := ToUnixPtr(null.Time{})
		if got != nil {
			t.Errorf("ToUnixPtr(invalid) = %v, want nil", got)
		}
	})

	t.Run("valid returns pointer to epoch seconds", func(t *testing.T) {
		tm := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		got := ToUnixPtr(null.TimeFrom(tm))
		if got == nil {
			t.Fatal("ToUnixPtr(valid) should not be nil")
		}
		if *got != float64(tm.Unix()) {
			t.Errorf("got = %v, want %v", *got, float64(tm.Unix()))
		}
	})
}

func TestRoundTripPair(t *testing.T) {
	original := 1735689600.0
	got := ToUnix(FromUnix(original))
	if got != original {
		t.Errorf("int round-trip: got %v, want %v", got, original)
	}

	if FromUnixPtr(ToUnixPtr(null.Time{Valid: false})) != (null.Time{}) {
		t.Errorf("invalid null.Time round-trip should yield invalid null.Time")
	}
	if ToUnixPtr(FromUnixPtr(nil)) != nil {
		t.Errorf("nil *float64 round-trip should yield nil")
	}
}

func TestHelperValuesAreLargeEnough(t *testing.T) {
	post2000 := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	if ToUnix(post2000) <= 0 {
		t.Errorf("post-2000 Unix seconds should be positive, got %v", ToUnix(post2000))
	}
}
