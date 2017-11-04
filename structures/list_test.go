package structures_test

import (
	"reflect"
	"testing"

	s "github.com/stsh89/algorithm/structures"
)

var listTests = []struct {
	numbers []int
	size    int
}{
	{[]int{}, 0},
	{[]int{0}, 1},
	{[]int{1, 2}, 2},
	{[]int{1, 2, 3, 4, 5, 6, 7}, 7},
}

func TestListSize(t *testing.T) {
	for _, tt := range listTests {
		list := s.NewList(tt.numbers)
		got := list.Size()
		want := tt.size

		if got != want {
			t.Errorf("list.Size(), got %v, want %v", got, want)
		}
	}
}

func TestListNumbers(t *testing.T) {
	for _, tt := range listTests {
		list := s.NewList(tt.numbers)
		got := list.Numbers()
		want := tt.numbers

		if !reflect.DeepEqual(got, want) {
			t.Errorf("list.Root(), got %v, want %v", got, want)
		}
	}
}
