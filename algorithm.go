package algorithm

import (
	"github.com/stsh89/algorithm/sorting"
)

func New(numbers []int) Algorithm {
	return Algorithm{numbers}
}

type Algorithm struct {
	numbers []int
}

func (a *Algorithm) sort(sortStrategy func([]int)) []int {
	tmp := append([]int(nil), a.numbers...)
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
	return a.sort(sorting.Insert)
}

func (a *Algorithm) BubbleSort() []int {
	return a.sort(sorting.Bubble)
}

func (a *Algorithm) MergeSort() []int {
	return a.sort(sorting.Merge)
}
