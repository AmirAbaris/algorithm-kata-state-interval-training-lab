package interval_overlaps_test

import (
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/interval_overlaps"
)

func TestOverlaps_ClearOverlap(t *testing.T) {
	if !interval_overlaps.Overlaps(kata.Interval{1, 3}, kata.Interval{2, 5}) {
		t.Fatal("expected overlap")
	}
}

func TestOverlaps_Separated(t *testing.T) {
	if interval_overlaps.Overlaps(kata.Interval{1, 3}, kata.Interval{4, 5}) {
		t.Fatal("expected no overlap")
	}
}

func TestOverlaps_Touching(t *testing.T) {
	if !interval_overlaps.Overlaps(kata.Interval{1, 3}, kata.Interval{3, 5}) {
		t.Fatal("touching intervals should overlap")
	}
}

func TestOverlaps_Contained(t *testing.T) {
	if !interval_overlaps.Overlaps(kata.Interval{1, 10}, kata.Interval{3, 5}) {
		t.Fatal("expected overlap")
	}
}

func TestOverlaps_SinglePointSame(t *testing.T) {
	if !interval_overlaps.Overlaps(kata.Interval{5, 5}, kata.Interval{5, 5}) {
		t.Fatal("expected overlap")
	}
}

func TestOverlaps_SinglePointSeparate(t *testing.T) {
	if interval_overlaps.Overlaps(kata.Interval{5, 5}, kata.Interval{6, 6}) {
		t.Fatal("expected no overlap")
	}
}

func TestOverlaps_AdjacentGap(t *testing.T) {
	if interval_overlaps.Overlaps(kata.Interval{1, 2}, kata.Interval{4, 5}) {
		t.Fatal("expected no overlap")
	}
}
