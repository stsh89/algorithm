package sortings_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
)

func TestRandomizedQuick(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.RandomizedQuick(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("RandomizedQuick(), got %v, want %v", tt.In, tt.Out)
		}
	}
}
