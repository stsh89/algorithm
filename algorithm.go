package algorithm

import (
	"github.com/stsh89/algorithm/problems"
	"github.com/stsh89/algorithm/random"
	"github.com/stsh89/algorithm/sortings"
)

func New(numbers []int) Algorithm {
	return Algorithm{numbers}
}

type Algorithm struct {
	numbers []int
}

func (a *Algorithm) sort(sortStrategy func([]int)) []int {
	tmp := make([]int, len(a.numbers))
	copy(tmp, a.numbers)
	sortStrategy(tmp)
	return tmp
}

func (a *Algorithm) GetNumbers() []int {
	return a.numbers
}

func (a *Algorithm) SetNumbers(numbers []int) {
	a.numbers = numbers
}

func (a *Algorithm) InsertSort() []int {
	return a.sort(sortings.Insert)
}

func (a *Algorithm) BubbleSort() []int {
	return a.sort(sortings.Bubble)
}

func (a *Algorithm) MergeSort() []int {
	return a.sort(sortings.Merge)
}

func (a *Algorithm) HeapSort() []int {
	return a.sort(sortings.Heap)
}

func (a *Algorithm) QuickSort() []int {
	return a.sort(sortings.Quick)
}

func (a *Algorithm) RandomizedQuickSort() []int {
	return a.sort(sortings.RandomizedQuick)
}

func (a *Algorithm) CountingSort(maxNumber int) []int {
	return sortings.Counting(a.numbers, maxNumber)
}

func (a *Algorithm) SubArrayProblem() int {
	return problems.SubArray(a.numbers)
}

func (a *Algorithm) PermuteBySorting() []int {
	return a.sort(random.PermuteBySorting)
}

func (a *Algorithm) RandomizeInPlace() []int {
	return a.sort(random.RandomizeInPlace)
}
