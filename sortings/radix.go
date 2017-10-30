//
// RADIX-SORT(A, d)
// 	for i = 1 to d
// 		use a stable sort to sort array A on digit i
//

package sortings

func Radix(numbers []int, d int) {
	if len(numbers) < 2 {
		return
	}

	for i := 1; i <= d; i++ {
		CountingByDigitNumber(numbers, 1000, i)
	}
}
