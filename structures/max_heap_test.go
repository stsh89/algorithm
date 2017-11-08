package structures_test

import (
	"reflect"
	"testing"

	"github.com/stsh89/algorithm/structures"
)

func TestGetNumbers0(t *testing.T) {
	in := []int{}
	heap := structures.NewMaxHeap(in)
	got := heap.GetNumbers()
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers(), got %v, want %v", got, want)
	}
}

func TestGetNumbers1(t *testing.T) {
	in := []int{1}
	heap := structures.NewMaxHeap(in)
	got := heap.GetNumbers()
	want := []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers(), got %v, want %v", got, want)
	}
}

func TestGetNumbers2(t *testing.T) {
	in := []int{2}
	heap := structures.NewMaxHeap(in)
	got := heap.GetNumbers()
	want := []int{2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers(), got %v, want %v", got, want)
	}
}

func TestGetNumbersMany(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	heap := structures.NewMaxHeap(in)
	got := heap.GetNumbers()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetNumbers(), got %v, want %v", got, want)
	}
}

func TestGetSize0(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.GetSize()
	want := 0

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}
}

func TestGetSize1(t *testing.T) {
	heap := structures.NewMaxHeap([]int{1})
	got := heap.GetSize()
	want := 1

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}
}

func TestGetSize2(t *testing.T) {
	heap := structures.NewMaxHeap([]int{1, 2})
	got := heap.GetSize()
	want := 2

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}
}

func TestGetSizeMany(t *testing.T) {
	heap := structures.NewMaxHeap([]int{1, 2, 3, 4, 5, 6, 7})
	got := heap.GetSize()
	want := 7

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}
}

func TestLeft1(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Left(1)
	want := 2

	if got != want {
		t.Errorf("Left(), got %v, want %v", got, want)
	}
}

func TestLeft2(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Left(2)
	want := 4

	if got != want {
		t.Errorf("Left(), got %v, want %v", got, want)
	}
}

func TestLeft3(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Left(3)
	want := 6

	if got != want {
		t.Errorf("Left(), got %v, want %v", got, want)
	}
}

func TestRight1(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Right(1)
	want := 3

	if got != want {
		t.Errorf("Right(), got %v, want %v", got, want)
	}
}

func TestRight2(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Right(2)
	want := 5

	if got != want {
		t.Errorf("Right(), got %v, want %v", got, want)
	}
}

func TestRight3(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.Right(3)
	want := 7

	if got != want {
		t.Errorf("Right(), got %v, want %v", got, want)
	}
}

func TestHeapifyMany(t *testing.T) {
	heap := structures.NewMaxHeap([]int{1, 2, 3, 4, 5, 6, 7})
	heap.Heapify(1)
	got := heap.GetNumbers()
	want := []int{3, 2, 7, 4, 5, 6, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Right(), got %v, want %v", got, want)
	}
}

func TestBuildMany(t *testing.T) {
	heap := structures.NewMaxHeap([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
	heap.Build()
	got := heap.GetNumbers()
	want := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Build(), got %v, want %v", got, want)
	}
}

func TestSetHeapSize(t *testing.T) {
	heap := structures.NewMaxHeap([]int{})
	got := heap.GetSize()
	want := 0

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}

	heap.SetHeapSize(10)
	got = heap.GetSize()
	want = 0

	if got != want {
		t.Errorf("GetSize(), got %v, want %v", got, want)
	}
}

func TestSwap(t *testing.T) {
	heap := structures.NewMaxHeap([]int{0, 1})
	heap.Swap(0, 1)
	got := heap.GetNumbers()
	want := []int{1, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Swap(), got %v, want %v", got, want)
	}
}
