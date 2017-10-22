package utils_test

import (
	"github.com/stsh89/algorithm/utils"
	"testing"
)

var Min = utils.Min

func MinTest1(t *testing.T) {
	a, b := 1, 3
	got := Min(a, b)
	want := 1

	if got != want {
		t.Errorf("Min(%d, %d), got %d, want %d", a, b, got, want)
	}
}

func MinTest2(t *testing.T) {
	a, b := 1, 1
	got := Min(a, b)
	want := 1

	if got != want {
		t.Errorf("Min(%d, %d), got %d, want %d", a, b, got, want)
	}
}
