package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

var sortingsTestSet = []struct {
	in  []int
	out []int
	d   int
}{
	{[]int{}, []int{}, 0},
	{[]int{1}, []int{1}, 1},
	{[]int{2, 1}, []int{1, 2}, 1},
	{[]int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}, 1},
	{[]int{17, 16, 15, 14, 13, 12, 11}, []int{11, 12, 13, 14, 15, 16, 17}, 2},
	{[]int{176, 165, 157, 141, 134, 123, 113}, []int{113, 123, 134, 141, 157, 165, 176}, 3},
}

func TestRadix(t *testing.T) {
	for _, tt := range sortingsTestSet {
		got := sortings.Radix(tt.in, tt.d)

		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("Radix(%v, %v), got %v, want %v", tt.in, tt.d, got, tt.out)
		}
	}
}
