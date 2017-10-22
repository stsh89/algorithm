package random

import (
	"crypto/rand"
	"math/big"
)

func random(min, max int) int64 {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return res.Int64() + int64(min)
}

func seedRandomNumbers(numbers []int64) {
	for i := 0; i < len(numbers); i++ {
		numbers[i] = random(1, len(numbers)*len(numbers)*len(numbers))
	}
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

	tmp := make([]int64, len(numbers))
	seedRandomNumbers(tmp)
	sortByRandomNumbers(numbers, tmp)
}
