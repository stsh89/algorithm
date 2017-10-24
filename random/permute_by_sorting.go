package random

import (
	"github.com/stsh89/algorithm/utils"
)

func seedRandomNumbers(numbers []int) []int64 {
	res := make([]int64, len(numbers))

	for i := 0; i < len(res); i++ {
		res[i] = utils.Random(1, len(res)*len(res)*len(res))
	}

	return res
}

func sortByRandomNumbers(a []int, b []int64) {
	for i := 1; i < len(b); i++ {
		for j := len(b) - 1; j >= i; j-- {
			if b[j] < b[j-1] {
				b[j], b[j-1] = b[j-1], b[j]
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func PermuteBySorting(numbers []int) {
	if len(numbers) < 2 {
		return
	}

	randomNumbers := seedRandomNumbers(numbers)
	sortByRandomNumbers(numbers, randomNumbers)
}
