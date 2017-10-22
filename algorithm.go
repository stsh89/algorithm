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
	return a.sort(sortings.Insert)
}

func (a *Algorithm) BubbleSort() []int {
	return a.sort(sortings.Bubble)
}

func (a *Algorithm) MergeSort() []int {
	return a.sort(sortings.Merge)
}

func (a *Algorithm) SubArrayProblem() int {
	return problems.SubArray(a.numbers)
}

func (a *Algorithm) PermuteBySorting() []int {
	return a.sort(random.PermuteBySorting)
}
