package timeutil

import (
	"fmt"
	"time"
)

// Package timeutil provides minimal, company-wide utilities for time handling.
// Input: time strings/time values + an explicit *time.Location.
// Output: parsed/normalized time values with predictable timezone behavior
// (no dependency on environment-specific time.Local).

// Layout constants are centralized here to avoid scattered "magic strings".
const (
	// DefaultTimeLayout
	// Input: a time.Time to format (or a string to parse with ParseTimeIn).
	// Output: "YYYY-MM-DD HH:mm:ss" layout (no timezone in the string).
	DefaultTimeLayout = "2006-01-02 15:04:05"

	// DateLayout
	// Input: a time.Time to format (or a string to parse with ParseDateIn).
	// Output: "YYYY-MM-DD" layout.
	DateLayout = "2006-01-02"

	// LogTimestampLayout
	// Input: a time.Time to format for structured logging.
	// Output: RFC3339-like timestamp layout with milliseconds and offset.
	LogTimestampLayout = "2006-01-02T15:04:05.999Z07:00"
)

// ParseTimeIn parses a "YYYY-MM-DD HH:mm:ss" string in the given location.
// Input: s (time string), loc (timezone to interpret the string in).
// Output: time.Time located in loc; error if loc is nil or parsing fails.
func ParseTimeIn(s string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, fmt.Errorf("location is nil")
	}
	t, err := time.ParseInLocation(DefaultTimeLayout, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time (%s) layout=%q: %w", s, DefaultTimeLayout, err)
	}
	return t, nil
}

// ParseDateIn parses a "YYYY-MM-DD" string in the given location.
// Input: s (date string), loc (timezone to interpret the string in).
// Output: time.Time located in loc; error if loc is nil or parsing fails.
func ParseDateIn(s string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, fmt.Errorf("location is nil")
	}
	t, err := time.ParseInLocation(DateLayout, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse date (%s) layout=%q: %w", s, DateLayout, err)
	}
	return t, nil
}

// MidnightIn returns the start of day (00:00:00.000) for t in the given location.
// Input: t (any time), loc (day boundary timezone).
// Output: time at 00:00:00.000 in loc for the same YYYY-MM-DD as t.In(loc);
// error if loc is nil.
func MidnightIn(t time.Time, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, fmt.Errorf("location is nil")
	}
	tt := t.In(loc)
	return time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, loc), nil
}

// TruncateToHourIn truncates t down to the hour in the given location.
// Input: t (any time), loc (hour boundary timezone).
// Output: time with minute/second/nanosecond set to 0 in loc for the same
// YYYY-MM-DD HH as t.In(loc); error if loc is nil.
func TruncateToHourIn(t time.Time, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Time{}, fmt.Errorf("location is nil")
	}
	tt := t.In(loc)
	return time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), 0, 0, 0, loc), nil
}

// IsValidDateRange checks if start is not after end.
// Input: start, end (time.Time).
// Output: true if start <= end, otherwise false.
func IsValidDateRange(start, end time.Time) bool {
	return !start.After(end)
}

// ToUTC converts t to UTC without changing the instant.
// Input: t (time.Time).
// Output: the same instant represented in UTC (t.UTC()).
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}
