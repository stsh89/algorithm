package problems

import (
	"github.com/stsh89/algorithm/utils"
)

func SubArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	current, max := numbers[0], numbers[0]

	for i := 1; i < len(numbers); i++ {
		current = utils.Max(numbers[i], current+numbers[i])
		max = utils.Max(max, current)
	}

	return max
}
