package sortings

import (
	"math"
)

// CountingByDigitNumber sorts slice of integers
func CountingByDigitNumber(numbers []int, maxNumber, digitNumber int) {
	if len(numbers) < 2 {
		return
	}

	res := make([]int, len(numbers))
	k := maxNumber + 1
	c := make([]int, k)

	for i := 0; i < k; i++ {
		c[i] = 0
	}

	for j := 0; j < len(numbers); j++ {
		c[digit(numbers[j], digitNumber)]++
	}

	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}

	for j := len(numbers) - 1; j >= 0; j-- {
		res[c[digit(numbers[j], digitNumber)]-1] = numbers[j]
		c[digit(numbers[j], digitNumber)]--
	}

	for j := 0; j < len(numbers); j++ {
		numbers[j] = res[j]
	}
}

func digit(number int, digitNumber int) int {
	res := (float64(number) / (math.Pow(float64(10), float64(digitNumber-1))))
	return int(res) % 10
}
