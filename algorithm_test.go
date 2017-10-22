package algorithm_test

import (
	"github.com/stsh89/algorithm"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.InsertSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestBubbleSort(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.BubbleSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestMergeSort(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.MergeSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestSubArrayProblem(t *testing.T) {
	algorithm := algorithm.New([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	got := algorithm.SubArrayProblem()
	want := 6

	if got != want {
		t.Errorf("SubArrayProblem() == %d, want %d", got, want)
	}
}

func TestPermuteBySorting(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.PermuteBySorting()
	want := []int{2, 4, 3, 7, 5, 6, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}
