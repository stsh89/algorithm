package sortings_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/sortings"
	"github.com/stsh89/algorithm/test_utils"
)

func TestBubble(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		sortings.Bubble(tt.In)

		if !reflect.DeepEqual(tt.In, tt.Out) {
			t.Errorf("Bubble(%v), got %v, want %v", tt.In, tt.In, tt.Out)
		}
	}
}
