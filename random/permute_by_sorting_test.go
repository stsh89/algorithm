package random_test

import (
	"github.com/stsh89/algorithm/random"
	"reflect"
	"testing"
)

func TestPermuteBySorting0(t *testing.T) {
	in := []int{}
	random.PermuteBySorting(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("permuteBySorting(%v), want %v", in, want)
	}
}

func TestPermuteBySorting1(t *testing.T) {
	in := []int{1}
	random.PermuteBySorting(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("permuteBySorting(%v), want %v", in, want)
	}
}

func TestPermuteBySortingMany(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	random.PermuteBySorting(in)
	dont_want := []int{1, 2, 3, 4, 5, 6, 7}

	if reflect.DeepEqual(in, dont_want) {
		t.Errorf("permuteBySorting(%v), want %v", in, dont_want)
	}
}
