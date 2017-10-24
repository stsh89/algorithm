package random_test

import (
	"github.com/stsh89/algorithm/random"
	"reflect"
	"testing"
)

var randomizeInPlace = random.RandomizeInPlace

func TestRandomizeInPlace0(t *testing.T) {
	in := []int{}
	randomizeInPlace(in)
	want := []int{}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("randomizeInPlace(%v), want %v", in, want)
	}
}

func TestRandomizeInPlace1(t *testing.T) {
	in := []int{1}
	randomizeInPlace(in)
	want := []int{1}

	if !reflect.DeepEqual(in, want) {
		t.Errorf("randomizeInPlace(%v), want %v", in, want)
	}
}

func TestRandomizeInPlaceMany(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	randomizeInPlace(in)
	dont_want := []int{1, 2, 3, 4, 5, 6, 7}

	if reflect.DeepEqual(in, dont_want) {
		t.Errorf("randomizeInPlace(%v), want %v", in, dont_want)
	}
}
