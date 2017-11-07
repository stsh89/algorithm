package structures_test

import (
	"reflect"
	"testing"

	s "github.com/stsh89/algorithm/structures"
)

var listTests = []struct {
	numbers       []int
	size          int
	sortedNUmbers []int
}{
	{[]int{}, 0, []int{}},
	{[]int{0}, 1, []int{0}},
	{[]int{2, 1}, 2, []int{1, 2}},
	{[]int{1, 5, 4}, 3, []int{1, 4, 5}},
	{[]int{7, 6, 5, 4, 3, 2, 1}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
	{[]int{12, 1, 5, 3, 7, 13, 2}, 7, []int{1, 2, 3, 5, 7, 12, 13}},
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

func TestListSort(t *testing.T) {
	for _, tt := range listTests {
		list := s.NewList(tt.numbers)
		list.Sort()
		got := list.Numbers()
		want := tt.sortedNUmbers

		if !reflect.DeepEqual(got, want) {
			t.Errorf("list.Root(), got %v, want %v", got, want)
		}
	}
}
