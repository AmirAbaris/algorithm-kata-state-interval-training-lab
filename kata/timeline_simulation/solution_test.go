package timeline_simulation_test

import (
	"reflect"
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/timeline_simulation"
)

func TestSimulate_Empty(t *testing.T) {
	got := timeline_simulation.Simulate(nil)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestSimulate_SingleEvent(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 3, End: 5, Delta: 4},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 3, End: 5}, State: 4},
	}
	assertSegments(t, got, want)
}

func TestSimulate_NonOverlapping(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 1, End: 2, Delta: 5},
		{Start: 5, End: 7, Delta: 3},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 2}, State: 5},
		{Interval: kata.Interval{Start: 5, End: 7}, State: 3},
	}
	assertSegments(t, got, want)
}

func TestSimulate_OverlappingStack(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 1, End: 3, Delta: 2},
		{Start: 2, End: 4, Delta: 3},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 1}, State: 2},
		{Interval: kata.Interval{Start: 2, End: 3}, State: 5},
		{Interval: kata.Interval{Start: 4, End: 4}, State: 3},
	}
	assertSegments(t, got, want)
}

func TestSimulate_SameStartDifferentEnd(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 5, End: 5, Delta: 10},
		{Start: 5, End: 7, Delta: -3},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 5, End: 5}, State: 7},
		{Interval: kata.Interval{Start: 6, End: 7}, State: -3},
	}
	assertSegments(t, got, want)
}

func TestSimulate_AdjacentSameStateMerged(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 1, End: 3, Delta: 2},
		{Start: 4, End: 6, Delta: 2},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 3}, State: 2},
		{Interval: kata.Interval{Start: 4, End: 6}, State: 2},
	}
	assertSegments(t, got, want)
}

func TestSimulate_NegativeDelta(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 0, End: 2, Delta: 10},
		{Start: 1, End: 1, Delta: -15},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 0, End: 0}, State: 10},
		{Interval: kata.Interval{Start: 1, End: 1}, State: -5},
		{Interval: kata.Interval{Start: 2, End: 2}, State: 10},
	}
	assertSegments(t, got, want)
}

func TestSimulate_FullOverlapNested(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 1, End: 10, Delta: 1},
		{Start: 3, End: 7, Delta: 9},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 2}, State: 1},
		{Interval: kata.Interval{Start: 3, End: 7}, State: 10},
		{Interval: kata.Interval{Start: 8, End: 10}, State: 1},
	}
	assertSegments(t, got, want)
}

func TestSimulate_UnsortedEvents(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 5, End: 6, Delta: 1},
		{Start: 1, End: 2, Delta: 2},
		{Start: 2, End: 3, Delta: 3},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 1}, State: 2},
		{Interval: kata.Interval{Start: 2, End: 3}, State: 5},
		{Interval: kata.Interval{Start: 5, End: 6}, State: 1},
	}
	assertSegments(t, got, want)
}

func TestSimulate_ThreeWayOverlap(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 1, End: 5, Delta: 1},
		{Start: 3, End: 7, Delta: 2},
		{Start: 5, End: 9, Delta: 4},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 1, End: 2}, State: 1},
		{Interval: kata.Interval{Start: 3, End: 4}, State: 3},
		{Interval: kata.Interval{Start: 5, End: 5}, State: 7},
		{Interval: kata.Interval{Start: 6, End: 7}, State: 6},
		{Interval: kata.Interval{Start: 8, End: 9}, State: 4},
	}
	assertSegments(t, got, want)
}

func TestSimulate_SinglePointEvent(t *testing.T) {
	got := timeline_simulation.Simulate([]kata.Event{
		{Start: 42, End: 42, Delta: 100},
	})
	want := []kata.StateSegment{
		{Interval: kata.Interval{Start: 42, End: 42}, State: 100},
	}
	assertSegments(t, got, want)
}

func assertSegments(t *testing.T, got, want []kata.StateSegment) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
