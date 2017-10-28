//
// COUNTING-SORT
// 	let C[0..k] be a new array
// 	for i = 0 to k
// 		C[i] = 0
// 	for j = 1 to A.length
// 		C[A[j]] = C[A[j]] + i, so C[i] contains number of elements equal to i
// 	for i = 1 to k
// 		C[i] = C[i] + C[i-1], so C[i] contains number of elements less or equal to i
// 	for j = A.length down to 1
// 		B[C[A[j]]] = A[j]
// 		C[A[j]] = C[A[j]] - 1
//

package sortings

func Counting(numbers []int, maxNumber int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	res := make([]int, len(numbers))
	k := maxNumber + 1
	c := make([]int, k)

	for i := 0; i < k; i++ {
		c[i] = 0
	}

	for j := 0; j < len(numbers); j++ {
		c[numbers[j]]++
	}

	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}

	for j := len(numbers) - 1; j >= 0; j-- {
		res[c[numbers[j]]-1] = numbers[j]
		c[numbers[j]]--
	}

	return res
}
