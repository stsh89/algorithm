package sortings

import (
	"github.com/stsh89/algorithm/structures"
)

func Heap(numbers []int) {
	sort(heap, numbers)
}

func heap(numbers []int) {
	heap := structures.NewMaxHeap(numbers)
	heap.Build()

	for i := len(numbers) - 1; i >= 1; i-- {
		heap.Swap(0, i)
		heap.SetHeapSize(heap.GetHeapSize() - 1)
		heap.Heapify(1)
	}
}
