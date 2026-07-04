package find_available_slots_test

import (
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/find_available_slots"
	"algorithm-kata/kata/testutil"
)

func TestFind_NoBusy(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{9, 17}},
		nil,
		2,
	)
	testutil.AssertIntervals(t, got, []kata.Interval{{9, 17}})
}

func TestFind_FullyBooked(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{9, 12}},
		[]kata.Interval{{9, 12}},
		1,
	)
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestFind_GapsWithMinDuration(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{9, 17}},
		[]kata.Interval{{11, 12}, {14, 15}},
		2,
	)
	testutil.AssertIntervals(t, got, []kata.Interval{{9, 10}, {16, 17}})
}

func TestFind_FiltersShortGaps(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{1, 10}},
		[]kata.Interval{{5, 5}},
		3,
	)
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 4}, {6, 10}})
}

func TestFind_MultipleWorkingBlocks(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{1, 5}, {8, 12}},
		[]kata.Interval{{3, 9}},
		2,
	)
	testutil.AssertIntervals(t, got, []kata.Interval{{1, 2}, {10, 12}})
}

func TestFind_MinDurationExactFit(t *testing.T) {
	got := find_available_slots.Find(
		[]kata.Interval{{10, 14}},
		[]kata.Interval{{12, 12}},
		2,
	)
	testutil.AssertIntervals(t, got, []kata.Interval{{10, 11}, {13, 14}})
}

func TestFind_EmptyWorkingHours(t *testing.T) {
	got := find_available_slots.Find(nil, []kata.Interval{{1, 5}}, 1)
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestDuration_Inclusive(t *testing.T) {
	if got := find_available_slots.Duration(kata.Interval{2, 5}); got != 4 {
		t.Fatalf("got %d, want 4", got)
	}
}
