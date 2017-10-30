package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

var countingTestSet = []struct {
	in          []int
	out         []int
	maxNumber   int
	digitNumber int
}{
	{[]int{}, []int{}, 0, 0},
	{[]int{1}, []int{1}, 1, 1},
	{[]int{2, 1}, []int{1, 2}, 2, 1},
	{[]int{20, 10}, []int{20, 10}, 0, 1},
	{[]int{20, 10}, []int{10, 20}, 2, 2},
	{[]int{70, 60, 50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50, 60, 70}, 7, 2},
	{[]int{127, 126, 125, 124, 123, 122, 121}, []int{121, 122, 123, 124, 125, 126, 127}, 7, 1},
	{[]int{171, 162, 153, 144, 135, 126, 117}, []int{117, 126, 135, 144, 153, 162, 171}, 7, 2},
	{[]int{727, 626, 525, 424, 323, 222, 121}, []int{121, 222, 323, 424, 525, 626, 727}, 7, 3},
}

func TestCountingByDigitNumber(t *testing.T) {
	for _, tt := range countingTestSet {
		sortings.CountingByDigitNumber(tt.in, tt.maxNumber, tt.digitNumber)

		if !reflect.DeepEqual(tt.in, tt.out) {
			t.Errorf("CountingByDigitNumber(%v, %v), got %v, want %v",
				tt.maxNumber, tt.digitNumber, tt.in, tt.out)
		}
	}
}
