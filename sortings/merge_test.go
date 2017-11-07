package sortings_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
)

func TestMerge(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.Merge(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("Merge(%v), got %v, want %v", tt.In, tt.In, tt.Out)
		}
	}
}
