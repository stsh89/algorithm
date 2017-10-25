package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

func TestHeapSort0(t *testing.T) {
	in := []int{}
	sortings.Heap(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("HeapSort(), got %v, want %v", in, want)
	}
}

func TestHeapSort1(t *testing.T) {
	in := []int{1}
	sortings.Heap(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
                t.Errorf("HeapSort(), got %v, want %v", in, want)
	}
}

func TestHeapSort2(t *testing.T) {
	in := []int{2, 1}
	sortings.Heap(in)
	want := []int{1, 2}

	if !reflect.DeepEqual(in, want) {
                t.Errorf("HeapSort(), got %v, want %v", in, want)
	}
}

func TestHeapSortMany(t *testing.T) {
	in := []int{7, 6, 5, 4, 3, 2, 1}
	sortings.Heap(in)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(in, want) {
                t.Errorf("HeapSort(), got %v, want %v", in, want)
	}
}
