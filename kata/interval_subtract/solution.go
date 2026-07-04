package interval_subtract

import (
	"algorithm-kata/kata"
	im "algorithm-kata/kata/interval_merge"
	"algorithm-kata/kata/split_interval"
)

func Subtract(base []kata.Interval, toRemove []kata.Interval) []kata.Interval {
	if len(base) == 0 {
		return nil
	}
	if len(toRemove) == 0 {
		return im.Merge(base)
	}

	fragments := im.Merge(base)
	removes := im.Merge(toRemove)

	for _, cut := range removes {
		var next []kata.Interval
		for _, f := range fragments {
			next = append(next, split_interval.Split(f, cut)...)
		}
		fragments = next
	}

	return im.Merge(fragments)
}
