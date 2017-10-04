package algorithm_test

import (
	"github.com/stsh89/algorithm"
	"reflect"
	"testing"
)

var ReverseInsertSort = algorithm.ReverseInsertSort

func TestReverseInsertSort0(t *testing.T) {
	in := []int{}
	got := ReverseInsertSort(in)
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestReverseInsertSort1(t *testing.T) {
	in := []int{1}
	got := ReverseInsertSort(in)
	want := []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestReverseInsertSort2(t *testing.T) {
	in := []int{1, 2}
	got := ReverseInsertSort(in)
	want := []int{2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestReverseInsertSortMany(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	got := ReverseInsertSort(in)
	want := []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}
