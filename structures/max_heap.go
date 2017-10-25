package structures

type MaxHeap struct {
	numbers []int
}

func (m *MaxHeap) GetNumbers() []int {
	return m.numbers
}

func (m *MaxHeap) GetSize() int {
	return len(m.numbers)
}

func NewMaxHeap(numbers []int) MaxHeap {
	return MaxHeap{numbers}
}

func (m *MaxHeap) Left(i int) int {
	return 2 * i
}

func (m *MaxHeap) Right(i int) int {
	return 2*i + 1
}

func (m *MaxHeap) Heapify(i int) {
	m.heapify(i - 1)
}

func (m *MaxHeap) Build() {
	for i := m.GetSize() / 2; i >= 1; i-- {
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
	left, right, largest := m.left(i), m.right(i), i

	if left < m.GetSize() && m.numbers[left] > m.numbers[i] {
		largest = left
	}

	if right < m.GetSize() && m.numbers[right] > m.numbers[largest] {
		largest = right
	}

	if largest != i {
		m.numbers[i], m.numbers[largest] = m.numbers[largest], m.numbers[i]
		m.Heapify(largest + 1)
	}
}
