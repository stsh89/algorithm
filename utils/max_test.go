package utils_test

import (
	"github.com/stsh89/algorithm/utils"
	"testing"
)

func TestMaxt1(t *testing.T) {
	a, b := 1, 3
	got := utils.Max(a, b)
	want := 3

	if got != want {
		t.Errorf("Max(%d, %d), got %d, want %d", a, b, got, want)
	}
}

func TestMax2(t *testing.T) {
	a, b := 1, 1
	got := utils.Max(a, b)
	want := 1

	if got != want {
		t.Errorf("Max(%d, %d), got %d, want %d", a, b, got, want)
	}
}
