package sum_numbers_test

import (
	"testing"

	"algorithm-kata/kata/sum_numbers"
)

func TestSum_Empty(t *testing.T) {
	if got := sum_numbers.Sum(nil); got != 0 {
		t.Fatalf("got %d, want 0", got)
	}
}

func TestSum_Single(t *testing.T) {
	if got := sum_numbers.Sum([]int{7}); got != 7 {
		t.Fatalf("got %d, want 7", got)
	}
}

func TestSum_Multiple(t *testing.T) {
	if got := sum_numbers.Sum([]int{3, 1, 4}); got != 8 {
		t.Fatalf("got %d, want 8", got)
	}
}

func TestSum_Negative(t *testing.T) {
	if got := sum_numbers.Sum([]int{10, -3, 2}); got != 9 {
		t.Fatalf("got %d, want 9", got)
	}
}

func TestSum_AllZeros(t *testing.T) {
	if got := sum_numbers.Sum([]int{0, 0, 0}); got != 0 {
		t.Fatalf("got %d, want 0", got)
	}
}
