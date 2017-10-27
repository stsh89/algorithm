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
		t.Errorf("BubbleSort() == %v, want %v", got, want)
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
		t.Errorf("MergeSort() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestHeapSort(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.HeapSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("HeapSort() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = []int{7, 6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}

func TestQuickSort(t *testing.T) {
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.QuickSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("QuickSort() == %v, want %v", got, want)
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
