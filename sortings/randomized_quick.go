//
// RANDOMIZED-QUICKSORT(A, p, r)
// 	if p < r
// 		q = RANDOMIZED-PARTITION(A, p, r)
// 		RANDOMIZED-QUICKSORT(A, p, q-1)
// 		RANDOMIZED-QUICKSORT(A, q+1, r)
//
// RANDOMIZED-PARTITION(A, p, r)
// 	i = RANDOM(p, r)
// 	exchange A[r] with A[i]
// 	return PARTITION(A, p, r)
//

package sortings

import (
	"github.com/stsh89/algorithm/utils"
)

// RandomizedQuick sorts slice of integers
func RandomizedQuick(numbers []int) {
	sort(randomizedQuick, numbers)
}

func randomizedQuick(numbers []int) {
	randomizedQuickSort(numbers, 0, len(numbers)-1)
}

func randomizedQuickSort(numbers []int, p, r int) {
	if p < r {
		q := randomizedPartition(numbers, p, r)
		randomizedQuickSort(numbers, p, q-1)
		randomizedQuickSort(numbers, q+1, r)
	}
}

func randomizedPartition(numbers []int, p, r int) int {
	randomIndex := utils.Random(p, r)
	numbers[randomIndex], numbers[r] = numbers[r], numbers[randomIndex]
	return partition(numbers, p, r)
}
