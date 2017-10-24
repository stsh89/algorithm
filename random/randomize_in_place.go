package random

import (
	"github.com/stsh89/algorithm/utils"
)

func RandomizeInPlace(numbers []int) {
	if len(numbers) < 2 {
		return
	}

	for i := 0; i < len(numbers); i++ {
		randomIndex := utils.Random(0, len(numbers)-1)
		numbers[i], numbers[randomIndex] = numbers[randomIndex], numbers[i]
	}
}
