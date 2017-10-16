package sortings

func Insert(numbers []int) {
	if len(numbers) < 2 {
		return
	}

	for i := 1; i < len(numbers); i++ {
		for j, curr := i, numbers[i]; (j > 0) && (numbers[j-1] > curr); j -= 1 {
			numbers[j] = numbers[j-1]
			numbers[j-1] = curr
		}
	}
}
