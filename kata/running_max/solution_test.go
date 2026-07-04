package running_max_test

import (
	"testing"

	"algorithm-kata/kata/running_max"
	"algorithm-kata/kata/testutil"
)

func TestRunningMax_Empty(t *testing.T) {
	got := running_max.RunningMax(nil)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestRunningMax_Single(t *testing.T) {
	testutil.AssertInts(t, running_max.RunningMax([]int{5}), []int{5})
}

func TestRunningMax_Increasing(t *testing.T) {
	testutil.AssertInts(t, running_max.RunningMax([]int{1, 2, 3}), []int{1, 2, 3})
}

func TestRunningMax_DropThenRise(t *testing.T) {
	testutil.AssertInts(t, running_max.RunningMax([]int{3, 1, 4, 2}), []int{3, 3, 4, 4})
}

func TestRunningMax_Plateau(t *testing.T) {
	testutil.AssertInts(t, running_max.RunningMax([]int{2, 2, 2}), []int{2, 2, 2})
}

func TestRunningMax_MultiplePeaks(t *testing.T) {
	testutil.AssertInts(t, running_max.RunningMax([]int{1, 5, 3, 5, 2}), []int{1, 5, 5, 5, 5})
}
