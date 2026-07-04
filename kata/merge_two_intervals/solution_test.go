package merge_two_intervals_test

import (
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/merge_two_intervals"
	"algorithm-kata/kata/testutil"
)

func TestMergeTwo_Overlap(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{1, 3}, kata.Interval{2, 5})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 5}})
}

func TestMergeTwo_Separate(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{1, 3}, kata.Interval{5, 7})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 3}, {5, 7}})
}

func TestMergeTwo_Touching(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{1, 4}, kata.Interval{4, 5})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 5}})
}

func TestMergeTwo_UnsortedArgs(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{5, 7}, kata.Interval{1, 3})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 3}, {5, 7}})
}

func TestMergeTwo_Contained(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{1, 10}, kata.Interval{3, 5})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 10}})
}

func TestMergeTwo_SinglePointsSeparate(t *testing.T) {
	got := merge_two_intervals.MergeTwo(kata.Interval{1, 1}, kata.Interval{3, 3})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 1}, {3, 3}})
}
