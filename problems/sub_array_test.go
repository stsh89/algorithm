package problems_test

import (
	"github.com/stsh89/algorithm/problems"
	"testing"
)

var SubArray = problems.SubArray

func TestSubArray0(t *testing.T) {
	in := []int{}
	got := SubArray(in)
	want := 0

	if got != want {
		t.Errorf("SubArray(%v), got %d, want %v", in, got, want)
	}
}

func TestSubArray1(t *testing.T) {
	in := []int{1}
	got := SubArray(in)
	want := 1

	if got != want {
		t.Errorf("SubArray(%v), got %d, want %v", in, got, want)
	}
}

func TestSubArray2(t *testing.T) {
	in := []int{1, 2}
	got := SubArray(in)
	want := 3

	if got != want {
		t.Errorf("SubArray(%v), got %d, want %v", in, got, want)
	}
}

func TestSubArrayMany(t *testing.T) {
	in := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	got := SubArray(in)
	want := 6

	if got != want {
		t.Errorf("SubArray(%v), got %d, want %v", in, got, want)
	}
}
