package testutil

import (
	"reflect"
	"testing"

	"algorithm-kata/kata"
)

func AssertInts(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func AssertIntervals(t *testing.T, got, want []kata.Interval) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func AssertSegments(t *testing.T, got, want []kata.StateSegment) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
