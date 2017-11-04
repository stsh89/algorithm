package structures

// MaxHeap representation of the MaxHeap
type MaxHeap struct {
	numbers []int
	size    int
}

// NewMaxHeap creates MaxHeap structure
func NewMaxHeap(numbers []int) MaxHeap {
	return MaxHeap{numbers, len(numbers)}
}

// GetNumbers returns numbers contained in a heap
func (m *MaxHeap) GetNumbers() []int {
	return m.numbers
}

// GetSize returns numbers of the numbers
func (m *MaxHeap) GetSize() int {
	return len(m.numbers)
}

// SetHeapSize sets a size of a heap
func (m *MaxHeap) SetHeapSize(size int) {
	m.size = size
}

// GetHeapSize returns a size of a heap
func (m *MaxHeap) GetHeapSize() int {
	return m.size
}

// Left returns left child of a number
func (m *MaxHeap) Left(i int) int {
	return 2 * i
}

// Right returns right child of a number
func (m *MaxHeap) Right(i int) int {
	return 2*i + 1
}

// Swap swaps to numbers
func (m *MaxHeap) Swap(i, j int) {
	m.numbers[i], m.numbers[j] = m.numbers[j], m.numbers[i]
}

// Heapify organizes numbers in a heap
func (m *MaxHeap) Heapify(i int) {
	m.heapify(i - 1)
}

// Build creaes MaxHeap structure
func (m *MaxHeap) Build() {
	for i := m.GetHeapSize() / 2; i >= 1; i-- {
		m.Heapify(i)
	}
}

func (m *MaxHeap) left(i int) int {
	return m.Left(i) + 1
}

func (m *MaxHeap) right(i int) int {
	return m.Right(i) + 1
}

func (m *MaxHeap) heapify(i int) {
	left, right, largest, size := m.left(i), m.right(i), i, m.GetHeapSize()

	if left < size && m.numbers[left] > m.numbers[i] {
		largest = left
	}

	if right < size && m.numbers[right] > m.numbers[largest] {
		largest = right
	}

	if largest != i {
		m.numbers[i], m.numbers[largest] = m.numbers[largest], m.numbers[i]
		m.Heapify(largest + 1)
	}
}
