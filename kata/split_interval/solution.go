package split_interval

import "algorithm-kata/kata"

// Split removes cut from base and returns the remaining interval fragments.
func Split(base, cut kata.Interval) []kata.Interval {
	lo := max(base.Start, cut.Start)
	hi := min(base.End, cut.End)
	if lo > hi {
		return []kata.Interval{base}
	}

	var out []kata.Interval
	if base.Start <= cut.Start-1 {
		out = append(out, kata.Interval{Start: base.Start, End: cut.Start - 1})
	}
	if cut.End+1 <= base.End {
		out = append(out, kata.Interval{Start: cut.End + 1, End: base.End})
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
