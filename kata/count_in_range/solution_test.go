package count_in_range_test

import (
	"testing"

	"algorithm-kata/kata/count_in_range"
)

func TestCount_Empty(t *testing.T) {
	if got := count_in_range.Count(nil, 1, 10); got != 0 {
		t.Fatalf("got %d, want 0", got)
	}
}

func TestCount_NoneInRange(t *testing.T) {
	if got := count_in_range.Count([]int{1, 10, 20}, 5, 8); got != 0 {
		t.Fatalf("got %d, want 0", got)
	}
}

func TestCount_SomeInRange(t *testing.T) {
	if got := count_in_range.Count([]int{1, 5, 3, 8, 2}, 2, 5); got != 3 {
		t.Fatalf("got %d, want 3", got)
	}
}

func TestCount_AllInRange(t *testing.T) {
	if got := count_in_range.Count([]int{3, 4, 5}, 1, 10); got != 3 {
		t.Fatalf("got %d, want 3", got)
	}
}

func TestCount_SinglePointRange(t *testing.T) {
	if got := count_in_range.Count([]int{4, 4, 5, 4}, 4, 4); got != 3 {
		t.Fatalf("got %d, want 3", got)
	}
}

func TestCount_BoundaryInclusive(t *testing.T) {
	if got := count_in_range.Count([]int{2, 5, 7}, 2, 7); got != 3 {
		t.Fatalf("got %d, want 3", got)
	}
}
