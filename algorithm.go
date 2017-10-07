package algorithm

import (
	"github.com/stsh89/algorithm/sorting"
)

type Algorithm struct {
	numbers []int
}

func New(numbers []int) Algorithm {
	return Algorithm{numbers}
}

func (a *Algorithm) GetNumbers() []int {
	return a.numbers
}

func (a *Algorithm) SetNumbers(numbers []int) {
	a.numbers = numbers
}

func (a *Algorithm) InsertSort() []int {
	return sorting.Insert(append([]int(nil), a.numbers...))
}
