package structures

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func NewNode(value int, next, prev *Node) *Node {
	return &Node{value, next, prev}
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) SetPrev(node *Node) {
	n.prev = node
}
