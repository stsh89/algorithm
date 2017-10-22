package random_test

import (
	"github.com/stsh89/algorithm/random"
	"reflect"
	"testing"
)

var permuteBySorting = random.PermuteBySorting

func TestPermuteBySorting0(t *testing.T) {
	in := []int{}
	permuteBySorting(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("permuteBySorting(%v), want %v", in, want)
	}
}

func TestPermuteBySorting1(t *testing.T) {
	in := []int{1}
	permuteBySorting(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("permuteBySorting(%v), want %v", in, want)
	}
}

func TestPermuteBySortingMany(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	permuteBySorting(in)
	want := []int{2, 4, 3, 7, 5, 6, 1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("permuteBySorting(%v), want %v", in, want)
	}
}
