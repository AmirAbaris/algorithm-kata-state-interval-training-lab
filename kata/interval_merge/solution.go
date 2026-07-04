package interval_merge

import (
	"algorithm-kata/kata"
	io "algorithm-kata/kata/interval_overlaps"
	sort "algorithm-kata/kata/sort_intervals"
)

func Merge(intervals []kata.Interval) []kata.Interval {
	if len(intervals) == 0 {
		return nil
	}

	sortesd := sort.SortByStart(intervals)
	out := []kata.Interval{sortesd[0]}

	for i := 1; i < len(sortesd); i++ {
		last := &out[len(out)-1]
		current := sortesd[i]

		if io.Overlaps(*last, current) {
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			out = append(out, current)
		}
	}

	return out
}
