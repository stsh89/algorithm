package sortings

// Quick sorts slice of integers
func Quick(numbers []int) {
	sort(quick, numbers)
}

func quick(numbers []int) {
	quickSort(numbers, 0, len(numbers)-1)
}

func quickSort(numbers []int, p, r int) {
	if p < r {
		q := partition(numbers, p, r)
		quickSort(numbers, p, q-1)
		quickSort(numbers, q+1, r)
	}
}

func partition(numbers []int, p, r int) int {
	x := numbers[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if numbers[j] <= x {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[r] = numbers[r], numbers[i+1]

	return i + 1
}
