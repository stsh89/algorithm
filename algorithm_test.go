package algorithm_test

import (
	"github.com/stsh89/algorithm"
	"github.com/stsh89/algorithm/test_utils"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.InsertSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.InsertSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

	}
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.BubbleSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.BubbleSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.MergeSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.MergeSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.HeapSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.HeapSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.QuickSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.QuickSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

	}
}

func TestRandomizedQuickSort(t *testing.T) {
	for _, tt := range test_utils.SortingsTestSet {
		algorithm := algorithm.New(tt.In)
		got := algorithm.RandomizedQuickSort()

		if !reflect.DeepEqual(got, tt.Out) {
			t.Errorf("algorithm.RandomizedQuickSort(), got %v, want %v", got, tt.Out)
		}

		got = algorithm.GetNumbers()

		if !reflect.DeepEqual(got, tt.In) {
			t.Errorf("GetNumbers() == %v, want %v", got, tt.In)
		}

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
	dont_want := []int{7, 6, 5, 4, 3, 2, 1}

	if reflect.DeepEqual(got, dont_want) {
		t.Errorf("PermuteBySorting() == %v, want %v", got, dont_want)
	}

	got = algorithm.GetNumbers()
	want := []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestRandomizeInPlace(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.RandomizeInPlace()
	dont_want := []int{7, 6, 5, 4, 3, 2, 1}

	if reflect.DeepEqual(got, dont_want) {
		t.Errorf("RandomizeInPlace() == %v, want %v", got, dont_want)
	}

	got = algorithm.GetNumbers()
	want := []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}
