package sortings_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
)

func TestInsert(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.Insert(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("Insert(%v), got %v, want %v", tt.In, tt.In, tt.Out)
		}
	}
}
