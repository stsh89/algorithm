package sortings_test

import (
	"github.com/stsh89/algorithm/sortings"
	"reflect"
	"testing"
)

func TestInsertSort0(t *testing.T) {
	in := []int{}
	sortings.Insert(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("InsertSort(%v), want %v", in, want)
	}
}

func TestInsertSort1(t *testing.T) {
	in := []int{1}
	sortings.Insert(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("InsertSort(%v), want %v", in, want)
	}
}

func TestInsertSort2(t *testing.T) {
	in := []int{2, 1}
	sortings.Insert(in)
	want := []int{1, 2}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("InsertSort(%v), want %v", in, want)
	}
}

func TestInsertSortMany(t *testing.T) {
	in := []int{7, 6, 5, 4, 3, 2, 1}
	sortings.Insert(in)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("InsertSort(%v), want %v", in, want)
	}
}
