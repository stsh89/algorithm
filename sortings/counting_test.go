package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

var countingByDigitNumberTestSet = []struct {
	in        []int
	out       []int
	maxNumber int
}{
	{[]int{}, []int{}, 0},
	{[]int{1}, []int{1}, 1},
	{[]int{2, 1}, []int{1, 2}, 2},
	{[]int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}, 7},
	{[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1}, 1},
}

func TestCounting(t *testing.T) {
	for _, tt := range countingByDigitNumberTestSet {
		got := sortings.Counting(tt.in, tt.maxNumber)

		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("Counting(), got %v, want %v", got, tt.out)
		}
	}
}
