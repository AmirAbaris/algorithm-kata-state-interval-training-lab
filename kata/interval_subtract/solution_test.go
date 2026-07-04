package interval_subtract_test

import (
	"reflect"
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/interval_subtract"
)

func TestSubtract_EmptyBase(t *testing.T) {
	got := interval_subtract.Subtract(nil, []kata.Interval{{Start: 1, End: 5}})
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestSubtract_EmptyRemove(t *testing.T) {
	base := []kata.Interval{{Start: 1, End: 5}, {Start: 8, End: 10}}
	got := interval_subtract.Subtract(base, nil)
	want := []kata.Interval{{Start: 1, End: 5}, {Start: 8, End: 10}}
	assertIntervals(t, got, want)
}

func TestSubtract_NoOverlap(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 3}},
		[]kata.Interval{{Start: 5, End: 7}},
	)
	want := []kata.Interval{{Start: 1, End: 3}}
	assertIntervals(t, got, want)
}

func TestSubtract_PartialOverlap(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 10}},
		[]kata.Interval{{Start: 3, End: 5}, {Start: 7, End: 8}},
	)
	want := []kata.Interval{
		{Start: 1, End: 2},
		{Start: 6, End: 6},
		{Start: 9, End: 10},
	}
	assertIntervals(t, got, want)
}

func TestSubtract_FullRemoval(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 5}, {Start: 8, End: 12}},
		[]kata.Interval{{Start: 0, End: 20}},
	)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestSubtract_SplitMiddle(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 10}},
		[]kata.Interval{{Start: 4, End: 6}},
	)
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 7, End: 10},
	}
	assertIntervals(t, got, want)
}

func TestSubtract_TouchingBoundary(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 10}},
		[]kata.Interval{{Start: 10, End: 12}},
	)
	want := []kata.Interval{{Start: 1, End: 9}}
	assertIntervals(t, got, want)
}

func TestSubtract_OverlappingBase(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 6}, {Start: 4, End: 10}},
		[]kata.Interval{{Start: 5, End: 7}},
	)
	want := []kata.Interval{
		{Start: 1, End: 4},
		{Start: 8, End: 10},
	}
	assertIntervals(t, got, want)
}

func TestSubtract_OverlappingRemove(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 1, End: 20}},
		[]kata.Interval{{Start: 2, End: 8}, {Start: 6, End: 12}},
	)
	want := []kata.Interval{
		{Start: 1, End: 1},
		{Start: 13, End: 20},
	}
	assertIntervals(t, got, want)
}

func TestSubtract_RemoveExactMatch(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 3, End: 7}},
		[]kata.Interval{{Start: 3, End: 7}},
	)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestSubtract_MultipleFragments(t *testing.T) {
	got := interval_subtract.Subtract(
		[]kata.Interval{{Start: 0, End: 100}},
		[]kata.Interval{{Start: 10, End: 20}, {Start: 30, End: 40}, {Start: 50, End: 60}},
	)
	want := []kata.Interval{
		{Start: 0, End: 9},
		{Start: 21, End: 29},
		{Start: 41, End: 49},
		{Start: 61, End: 100},
	}
	assertIntervals(t, got, want)
}

func assertIntervals(t *testing.T, got, want []kata.Interval) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
