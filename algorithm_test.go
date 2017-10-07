package algorithm_test

import (
	"github.com/stsh89/algorithm"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	numbers := []int{7, 6, 5, 4, 3, 2, 1}
	algorithm := algorithm.New([]int{7, 6, 5, 4, 3, 2, 1})
	got := algorithm.InsertSort()
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Insert() == %v, want %v", got, want)
	}

	got = algorithm.GetNumbers()
	want = numbers

	if !reflect.DeepEqual(got, numbers) {
		t.Errorf("GetNumbers() == %v, want %v", got, want)
	}
}
