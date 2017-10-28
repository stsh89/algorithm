package sortings

func Insert(numbers []int) {
	sort(insert, numbers)
}

func insert(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		for j, curr := i, numbers[i]; (j > 0) && (numbers[j-1] > curr); j-- {
			numbers[j] = numbers[j-1]
			numbers[j-1] = curr
		}
	}
}
