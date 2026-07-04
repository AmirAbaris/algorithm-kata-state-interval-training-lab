package sort_intervals_test

import (
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/sort_intervals"
	"algorithm-kata/kata/testutil"
)

func TestSortByStart_Empty(t *testing.T) {
	got := sort_intervals.SortByStart(nil)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestSortByStart_Single(t *testing.T) {
	testutil.AssertIntervals(t, sort_intervals.SortByStart([]kata.Interval{{2, 5}}), []kata.Interval{{2, 5}})
}

func TestSortByStart_Unsorted(t *testing.T) {
	got := sort_intervals.SortByStart([]kata.Interval{
		{Start: 8, End: 10},
		{Start: 1, End: 3},
		{Start: 2, End: 6},
	})
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 2, End: 6},
		{Start: 8, End: 10},
	}
	testutil.AssertIntervals(t, got, want)
}

func TestSortByStart_TieBreakByEnd(t *testing.T) {
	got := sort_intervals.SortByStart([]kata.Interval{
		{Start: 1, End: 5},
		{Start: 1, End: 2},
	})
	want := []kata.Interval{
		{Start: 1, End: 2},
		{Start: 1, End: 5},
	}
	testutil.AssertIntervals(t, got, want)
}

func TestSortByStart_DoesNotMerge(t *testing.T) {
	got := sort_intervals.SortByStart([]kata.Interval{
		{Start: 1, End: 4},
		{Start: 4, End: 5},
	})
	want := []kata.Interval{
		{Start: 1, End: 4},
		{Start: 4, End: 5},
	}
	testutil.AssertIntervals(t, got, want)
}
