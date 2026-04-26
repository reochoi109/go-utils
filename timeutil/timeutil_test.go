package timeutil

import (
	"testing"
	"time"
)

func TestParseTimeIn(t *testing.T) {
	loc := time.FixedZone("T", 9*60*60)

	got, err := ParseTimeIn("2026-04-26 10:20:30", loc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got.Location() != loc {
		t.Fatalf("expected location %v, got %v", loc, got.Location())
	}
	if got.Format(DefaultTimeLayout) != "2026-04-26 10:20:30" {
		t.Fatalf("unexpected parsed time: %s", got.Format(DefaultTimeLayout))
	}
}

func TestParseDateIn(t *testing.T) {
	loc := time.FixedZone("T", 9*60*60)

	got, err := ParseDateIn("2026-04-26", loc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got.Location() != loc {
		t.Fatalf("expected location %v, got %v", loc, got.Location())
	}
	if got.Format(DateLayout) != "2026-04-26" {
		t.Fatalf("unexpected parsed date: %s", got.Format(DateLayout))
	}
}

func TestMidnightIn(t *testing.T) {
	loc := time.FixedZone("T", 9*60*60)
	org := time.Date(2026, 4, 26, 19, 30, 33, 123, loc)

	got, err := MidnightIn(org, loc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got.Hour() != 0 || got.Minute() != 0 || got.Second() != 0 || got.Nanosecond() != 0 {
		t.Fatalf("expected midnight, got %s", got.Format(time.RFC3339Nano))
	}
}

func TestTruncateToHourIn(t *testing.T) {
	loc := time.FixedZone("T", 9*60*60)
	org := time.Date(2026, 4, 26, 19, 30, 33, 123, loc)

	got, err := TruncateToHourIn(org, loc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got.Hour() != 19 || got.Minute() != 0 || got.Second() != 0 || got.Nanosecond() != 0 {
		t.Fatalf("expected hour-truncated, got %s", got.Format(time.RFC3339Nano))
	}
}

func TestIsValidDateRange(t *testing.T) {
	a := time.Date(2026, 4, 26, 0, 0, 0, 0, time.UTC)
	b := time.Date(2026, 4, 27, 0, 0, 0, 0, time.UTC)

	if !IsValidDateRange(a, b) {
		t.Fatal("expected valid range")
	}
	if IsValidDateRange(b, a) {
		t.Fatal("expected invalid range")
	}
}

func TestToUTC(t *testing.T) {
	loc := time.FixedZone("T", 9*60*60)
	org := time.Date(2026, 4, 26, 10, 0, 0, 0, loc)

	got := ToUTC(org)
	if got.Location() != time.UTC {
		t.Fatalf("expected UTC, got %v", got.Location())
	}
	if !got.Equal(org) {
		t.Fatal("expected same instant")
	}
}
