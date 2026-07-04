package sort_intervals

import "algorithm-kata/kata"

// SortByStart returns intervals sorted by Start ascending, then End ascending.
func SortByStart(intervals []kata.Interval) []kata.Interval {
	// loop to loop the slices
	// compare start with end
	// create a new var to hold the smaller item
	// if we hit bigger, replace with smaller

	if len(intervals) == 0 {
		return nil
	}
	
	for i := range intervals {
		for j := 0; j < len(intervals)-1-i; j++ {
			if intervals[j].Start > intervals[j+1].Start {
				swap(&intervals[j], &intervals[j+1])
			}
			if intervals[j].Start == intervals[j+1].Start {
				if intervals[j].End > intervals[j+1].End {
					swap(&intervals[j], &intervals[j+1])
				}
			}
		}
	}

	return intervals
}

func swap(firstInterval, secInterval *kata.Interval) {
	temp := *firstInterval
	*firstInterval = *secInterval
	*secInterval = temp
}
