package utils_test

import (
	"github.com/stsh89/algorithm/utils"
	"testing"
)

func TestRandom1(t *testing.T) {
	a, b := 0, 1
	got := utils.Random(a, b)
	want := 0

	if got != int64(want) {
		t.Errorf("Min(%d, %d), got %d, want %d", a, b, got, want)
	}
}

func TestRandom2(t *testing.T) {
	a, b := 0, 2
	got := utils.Random(a, b)

	if got >= int64(b) || got < int64(a) {
		t.Errorf("Min(%d, %d), got %d, want to be less than %d", a, b, got, b)
	}
}

func TestRandom3(t *testing.T) {
	a, b := 0, 10
	got := utils.Random(a, b)

	if got >= int64(b) || got < int64(a) {
		t.Errorf("Min(%d, %d), got %d, want to be less than %d", a, b, got, b)
	}
}
