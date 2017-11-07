package structures

// List representation of a llinked list
type List struct {
	root *Node
	size int
}

// NewList return list structure
func NewList(numbers []int) *List {
	len := len(numbers)

	if len == 0 {
		return &List{root: nil, size: 0}
	}

	root := &Node{Value: numbers[0]}
	list := List{root: root, size: len}
	tmp := list.root

	for i := 1; i < len; i++ {
		tmp.Next = &Node{Value: numbers[i]}
		tmp = tmp.Next
	}

	return &list
}

// Size returns number of elements in a list
func (l *List) Size() int {
	return l.size
}

// Root return root node of the list
func (l *List) Root() *Node {
	return l.root
}

// Numbers returns numbers in a list
func (l *List) Numbers() []int {
	res := []int{}
	tmp := l.root

	for tmp != nil {
		res = append(res, tmp.Value)
		tmp = tmp.Next
	}

	return res
}

// Sort sorts list numbers
func (l *List) Sort() {
	if l.root == nil || l.root.Next == nil {
		return
	}

	curr, next := l.root, l.root.Next

	for next != nil {
		for curr != next {
			if curr.Value > next.Value {
				curr.Value, next.Value = next.Value, curr.Value
			}

			curr = curr.Next
		}

		curr = l.root
		next = next.Next
	}
}
