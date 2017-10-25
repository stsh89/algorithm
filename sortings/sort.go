package sortings

func sort(sortStrategy func([]int), numbers []int) {
	if len(numbers) < 2 {
		return
	}

	sortStrategy(numbers)
}
