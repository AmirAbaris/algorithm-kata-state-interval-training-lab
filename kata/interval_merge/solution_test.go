package interval_merge_test

import (
	"reflect"
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/interval_merge"
)

func TestMerge_Empty(t *testing.T) {
	got := interval_merge.Merge(nil)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestMerge_Single(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{{Start: 2, End: 5}})
	want := []kata.Interval{{Start: 2, End: 5}}
	assertIntervals(t, got, want)
}

func TestMerge_NonOverlapping(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 1, End: 2},
		{Start: 4, End: 5},
		{Start: 7, End: 9},
	})
	want := []kata.Interval{
		{Start: 1, End: 2},
		{Start: 4, End: 5},
		{Start: 7, End: 9},
	}
	assertIntervals(t, got, want)
}

func TestMerge_Overlapping(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 1, End: 3},
		{Start: 2, End: 6},
		{Start: 8, End: 10},
		{Start: 15, End: 18},
	})
	want := []kata.Interval{
		{Start: 1, End: 6},
		{Start: 8, End: 10},
		{Start: 15, End: 18},
	}
	assertIntervals(t, got, want)
}

func TestMerge_Touching(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 1, End: 4},
		{Start: 4, End: 5},
	})
	want := []kata.Interval{{Start: 1, End: 5}}
	assertIntervals(t, got, want)
}

func TestMerge_FullContainment(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 1, End: 10},
		{Start: 3, End: 5},
		{Start: 2, End: 7},
	})
	want := []kata.Interval{{Start: 1, End: 10}}
	assertIntervals(t, got, want)
}

func TestMerge_UnsortedInput(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 8, End: 10},
		{Start: 1, End: 3},
		{Start: 15, End: 18},
		{Start: 2, End: 6},
	})
	want := []kata.Interval{
		{Start: 1, End: 6},
		{Start: 8, End: 10},
		{Start: 15, End: 18},
	}
	assertIntervals(t, got, want)
}

func TestMerge_ChainTouching(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 1, End: 2},
		{Start: 2, End: 3},
		{Start: 3, End: 4},
	})
	want := []kata.Interval{{Start: 1, End: 4}}
	assertIntervals(t, got, want)
}

func TestMerge_DuplicateIntervals(t *testing.T) {
	got := interval_merge.Merge([]kata.Interval{
		{Start: 5, End: 8},
		{Start: 5, End: 8},
		{Start: 6, End: 7},
	})
	want := []kata.Interval{{Start: 5, End: 8}}
	assertIntervals(t, got, want)
}

func assertIntervals(t *testing.T, got, want []kata.Interval) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
