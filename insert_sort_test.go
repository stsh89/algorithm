package algorithm_test

import (
	"github.com/stsh89/algorithm"
	"reflect"
	"testing"
)

var InsertSort = algorithm.InsertSort

func TestInsertSort0(t *testing.T) {
	in := []int{}
	got := InsertSort(in)
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestInsertSort1(t *testing.T) {
	in := []int{1}
	got := InsertSort(in)
	want := []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestInsertSort2(t *testing.T) {
	in := []int{2, 1}
	got := InsertSort(in)
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}

func TestInsertSortMany(t *testing.T) {
	in := []int{7, 6, 5, 4, 3, 2, 1}
	got := InsertSort(in)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertSort(%v) == %v, want %v", in, got, want)
	}
}
