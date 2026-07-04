package find_available_slots

import "algorithm-kata/kata"

// Duration returns inclusive length of an interval.
func Duration(iv kata.Interval) int {
	return iv.End - iv.Start + 1
}

// Find returns free slots within workingHours after removing busy blocks,
// keeping only intervals at least minDuration long.
func Find(workingHours []kata.Interval, busy []kata.Interval, minDuration int) []kata.Interval {
	// TODO: implement
	return nil
}
