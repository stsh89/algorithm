package structures_test

import (
	"reflect"
	"testing"

	s "github.com/stsh89/algorithm/structures"
)

func TestNodeValueGet(t *testing.T) {
	node := s.Node{Value: 0}
	got := node.Value
	want := 0

	if got != want {
		t.Errorf("get node.Value, %v, %v", got, want)
	}
}

func TestNodeValueSet(t *testing.T) {
	node := s.Node{Value: 0}
	node.Value = 1
	got := node.Value
	want := 1

	if got != want {
		t.Errorf("set node.Value, %v, %v", got, want)
	}
}

func TestNodeNextGet(t *testing.T) {
	node := s.Node{Next: nil}
	got := node.Next
	var want *s.Node = nil

	if got != want {
		t.Errorf("get node.Next, %v, %v", got, want)
	}
}

func TestNodeNextSet(t *testing.T) {
	node := s.Node{Next: nil}
	node.Next = &s.Node{}
	got := node.Next
	want := &s.Node{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("set node.Next, %v, %v", got, want)
	}
}

func TestNodePrevGet(t *testing.T) {
	node := s.Node{Prev: nil}
	got := node.Prev
	var want *s.Node = nil

	if got != want {
		t.Errorf("get node.Prev, %v, %v", got, want)
	}
}

func TestNodePrevSet(t *testing.T) {
	node := s.Node{Prev: nil}
	node.Prev = &s.Node{}
	got := node.Prev
	want := &s.Node{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("set node.Prev, %v, %v", got, want)
	}
}
