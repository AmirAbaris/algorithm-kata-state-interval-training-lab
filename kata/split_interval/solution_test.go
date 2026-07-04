package split_interval_test

import (
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/split_interval"
	"algorithm-kata/kata/testutil"
)

func TestSplit_NoOverlap(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 5}, kata.Interval{6, 8})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 5}})
}

func TestSplit_MiddleRemoved(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 10}, kata.Interval{3, 5})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 2}, {6, 10}})
}

func TestSplit_FullRemoval(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 5}, kata.Interval{1, 5})
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestSplit_CutCoversBase(t *testing.T) {
	got := split_interval.Split(kata.Interval{2, 4}, kata.Interval{1, 10})
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestSplit_TrimLeft(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 10}, kata.Interval{1, 3})
	testutil.AssertIntervals(t, got, []kata.Interval{{4, 10}})
}

func TestSplit_TrimRight(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 10}, kata.Interval{8, 10})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 7}})
}

func TestSplit_CutInsideTouchingEdges(t *testing.T) {
	got := split_interval.Split(kata.Interval{1, 5}, kata.Interval{2, 4})
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 1}, {5, 5}})
}

func TestSplit_SinglePointBaseRemoved(t *testing.T) {
	got := split_interval.Split(kata.Interval{5, 5}, kata.Interval{5, 5})
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}
