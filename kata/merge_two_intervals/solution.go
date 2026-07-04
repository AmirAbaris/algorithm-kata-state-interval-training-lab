package merge_two_intervals

import (
	"algorithm-kata/kata"
	overlap "algorithm-kata/kata/interval_overlaps"
	sort "algorithm-kata/kata/sort_intervals"
)

// MergeTwo merges a and b if they overlap or touch; otherwise returns both sorted.
func MergeTwo(a, b kata.Interval) []kata.Interval {

	// check if overlaps
	// if yes merge in one interval -> [min start, max end]
	// if not sort and return

	overlaps := overlap.Overlaps(a, b)
	min := 0
	max := 0

	if overlaps {
		if a.Start > b.Start {
			min = b.Start
		} else {
			min = a.Start         
		}

		if a.End > b.End {
			max = a.End
		} else {
			max = b.End
		}

		return []kata.Interval{{Start: min, End: max}}
	}

	intervals := []kata.Interval{a, b}

	sorted := sort.SortByStart(intervals)

	return sorted
}
