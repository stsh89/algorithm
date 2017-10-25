package structures

type MaxHeap struct {
	numbers []int
        size int
}

func NewMaxHeap(numbers []int) MaxHeap {
	return MaxHeap{numbers, len(numbers)}
}

func (m *MaxHeap) GetNumbers() []int {
	return m.numbers
}

func (m *MaxHeap) GetSize() int {
	return len(m.numbers)
}

func (m *MaxHeap) SetHeapSize(size int) {
        m.size = size
}

func (m *MaxHeap) GetHeapSize() int {
        return m.size
}

func (m *MaxHeap) Left(i int) int {
	return 2 * i
}

func (m *MaxHeap) Right(i int) int {
	return 2*i + 1
}

func (m *MaxHeap) Swap(i, j int) {
	m.numbers[i], m.numbers[j] = m.numbers[j], m.numbers[i]
}

func (m *MaxHeap) Heapify(i int) {
	m.heapify(i - 1)
}

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
