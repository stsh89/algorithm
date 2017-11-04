package structures_test

import (
	s "github.com/stsh89/algorithm/structures"
	"reflect"
	"testing"
)

var testNodes = []struct {
	node     *s.Node
	nextNode *s.Node
	prevNode *s.Node
}{
	{s.NewNode(0, nil, nil),
		nil, nil},
	{s.NewNode(0, s.NewNode(1, nil, nil), s.NewNode(2, nil, nil)),
		s.NewNode(1, nil, nil), s.NewNode(2, nil, nil)},
}

func TestNodeGetValue(t *testing.T) {
	node := s.NewNode(0, nil, nil)
	got := node.GetValue()
	want := 0

	if got != want {
		t.Errorf("node.GetValue(), got %v, want %v", got, want)
	}
}

func TestNodeSetValue(t *testing.T) {
	node := s.NewNode(0, nil, nil)
	node.SetValue(1)
	got := node.GetValue()
	want := 1

	if got != want {
		t.Errorf("node.SetValue(), got %v, want %v", got, want)
	}
}

func TestNodeGetNext(t *testing.T) {
	for _, tt := range testNodes {
		got := tt.node.GetNext()
		want := tt.nextNode

		if !reflect.DeepEqual(got, want) {
			t.Errorf("node.GetNext(), got %v, want %v", got, want)
		}
	}
}

func TestNodeSetNext(t *testing.T) {
	node := s.NewNode(0, nil, nil)
	node.SetNext(s.NewNode(1, nil, nil))
	got := node.GetNext()
	want := s.NewNode(1, nil, nil)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("node.SetNext(), got %v, want %v", got, want)
	}
}

func TestNodeGetPrev(t *testing.T) {
	for _, tt := range testNodes {
		got := tt.node.GetPrev()
		want := tt.prevNode

		if !reflect.DeepEqual(got, want) {
			t.Errorf("node.GetPrev(), got %v, want %v", got, want)
		}
	}
}

func TestNodeSetPrev(t *testing.T) {
	node := s.NewNode(0, nil, nil)
	node.SetPrev(s.NewNode(1, nil, nil))
	got := node.GetPrev()
	want := s.NewNode(1, nil, nil)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("node.SetPrev(), got %v, want %v", got, want)
	}
}
