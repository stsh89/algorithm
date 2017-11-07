package sortings_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/sortings"
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
	{[]int{176, 165, 157, 141, 134, 122, 113}, []int{113, 122, 134, 141, 157, 165, 176}, 2},
	{[]int{176, 165, 157, 141, 134, 122, 113}, []int{141, 122, 113, 134, 165, 176, 157}, 1},
}

func TestRadix(t *testing.T) {
	for _, tt := range sortingsTestSet {
		sortings.Radix(tt.in, tt.d)

		if !reflect.DeepEqual(tt.in, tt.out) {
			t.Errorf("Radix(%v), got %v, want %v", tt.d, tt.in, tt.out)
		}
	}
}
