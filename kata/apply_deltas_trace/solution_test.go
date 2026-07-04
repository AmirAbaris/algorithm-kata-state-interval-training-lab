package apply_deltas_trace_test

import (
	"testing"

	"algorithm-kata/kata/apply_deltas_trace"
	"algorithm-kata/kata/testutil"
)

func TestTrace_EmptyDeltas(t *testing.T) {
	got := apply_deltas_trace.Trace(10, nil)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestTrace_SingleDelta(t *testing.T) {
	testutil.AssertInts(t, apply_deltas_trace.Trace(0, []int{5}), []int{5})
}

func TestTrace_Multiple(t *testing.T) {
	testutil.AssertInts(t, apply_deltas_trace.Trace(10, []int{3, -5, 2}), []int{13, 8, 10})
}

func TestTrace_NegativeStart(t *testing.T) {
	testutil.AssertInts(t, apply_deltas_trace.Trace(-2, []int{5, -1}), []int{3, 2})
}

func TestTrace_AllNegativeDeltas(t *testing.T) {
	testutil.AssertInts(t, apply_deltas_trace.Trace(100, []int{-30, -30}), []int{70, 40})
}
