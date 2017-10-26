package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.Heap(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("Heap(%v), got %v, want %v", tt.In, tt.In, tt.Out)
		}
	}
}
