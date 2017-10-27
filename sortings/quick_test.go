package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.Quick(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("Quick(%v), got %v, want %v", tt.In, tt.In, tt.Out)
		}
	}
}
