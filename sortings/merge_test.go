package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

var MergeSort = sortings.Merge

func TestMergeSort0(t *testing.T) {
	in := []int{}
	MergeSort(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("MergeSort(%v), want %v", in, want)
	}
}

func TestMergeSort1(t *testing.T) {
	in := []int{1}
	MergeSort(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("MergeSort(%v), want %v", in, want)
	}
}

func TestMergeSort2(t *testing.T) {
	in := []int{2, 1}
	MergeSort(in)
	want := []int{1, 2}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("MergeSort(%v), want %v", in, want)
	}
}

func TestMergeSortMany(t *testing.T) {
	in := []int{7, 6, 5, 4, 3, 2, 1}
	MergeSort(in)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("MergeSort(%v), want %v", in, want)
	}
}
