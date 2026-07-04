package interval_insert_test

import (
	"reflect"
	"testing"

	"algorithm-kata/kata"
	"algorithm-kata/kata/interval_insert"
)

func TestInsert_Empty(t *testing.T) {
	got := interval_insert.Insert(nil, kata.Interval{Start: 2, End: 5})
	want := []kata.Interval{{Start: 2, End: 5}}
	assertIntervals(t, got, want)
}

func TestInsert_NonOverlappingBefore(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 5, End: 7}, {Start: 10, End: 12}},
		kata.Interval{Start: 1, End: 3},
	)
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 5, End: 7},
		{Start: 10, End: 12},
	}
	assertIntervals(t, got, want)
}

func TestInsert_NonOverlappingAfter(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 3}, {Start: 5, End: 7}},
		kata.Interval{Start: 10, End: 12},
	)
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 5, End: 7},
		{Start: 10, End: 12},
	}
	assertIntervals(t, got, want)
}

func TestInsert_NonOverlappingMiddle(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 2}, {Start: 6, End: 7}},
		kata.Interval{Start: 3, End: 5},
	)
	want := []kata.Interval{
		{Start: 1, End: 2},
		{Start: 3, End: 5},
		{Start: 6, End: 7},
	}
	assertIntervals(t, got, want)
}

func TestInsert_OverlapLeft(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 3}, {Start: 6, End: 9}},
		kata.Interval{Start: 2, End: 5},
	)
	want := []kata.Interval{
		{Start: 1, End: 5},
		{Start: 6, End: 9},
	}
	assertIntervals(t, got, want)
}

func TestInsert_OverlapRight(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 3}, {Start: 6, End: 9}},
		kata.Interval{Start: 5, End: 7},
	)
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 5, End: 9},
	}
	assertIntervals(t, got, want)
}

func TestInsert_MergeBothSides(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 2}, {Start: 3, End: 5}, {Start: 6, End: 7}, {Start: 8, End: 10}, {Start: 12, End: 16}},
		kata.Interval{Start: 4, End: 8},
	)
	want := []kata.Interval{
		{Start: 1, End: 2},
		{Start: 3, End: 10},
		{Start: 12, End: 16},
	}
	assertIntervals(t, got, want)
}

func TestInsert_CoversAll(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 2, End: 4}, {Start: 6, End: 8}},
		kata.Interval{Start: 1, End: 10},
	)
	want := []kata.Interval{{Start: 1, End: 10}}
	assertIntervals(t, got, want)
}

func TestInsert_TouchingLeft(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 5, End: 7}, {Start: 10, End: 12}},
		kata.Interval{Start: 3, End: 5},
	)
	want := []kata.Interval{
		{Start: 3, End: 7},
		{Start: 10, End: 12},
	}
	assertIntervals(t, got, want)
}

func TestInsert_TouchingRight(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 3}, {Start: 7, End: 9}},
		kata.Interval{Start: 3, End: 7},
	)
	want := []kata.Interval{{Start: 1, End: 9}}
	assertIntervals(t, got, want)
}

func TestInsert_SinglePoint(t *testing.T) {
	got := interval_insert.Insert(
		[]kata.Interval{{Start: 1, End: 3}, {Start: 5, End: 7}},
		kata.Interval{Start: 4, End: 4},
	)
	want := []kata.Interval{
		{Start: 1, End: 3},
		{Start: 4, End: 4},
		{Start: 5, End: 7},
	}
	assertIntervals(t, got, want)
}

func assertIntervals(t *testing.T, got, want []kata.Interval) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
