package interval_overlaps

import "algorithm-kata/kata"

// Overlaps reports whether a and b overlap or touch.
func Overlaps(a, b kata.Interval) bool {
	// compare matching a and b and return result
	// compare a.End with b.Start
	// if b.Start <= a.End return true

	if a == b {
		return true
	}
	if a.Start <= b.End && b.Start <= a.End {
		return true
	}

	return false
}
