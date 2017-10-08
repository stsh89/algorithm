package sorting

func Insert(numbers []int) {
	if len(numbers) < 2 {
		return
	}

	for i := 1; i < len(numbers); i++ {
		for j, tmp := i, numbers[i]; (j > 0) && (numbers[j-1] > tmp); j -= 1 {
			numbers[j] = numbers[j-1]
			numbers[j-1] = tmp
		}
	}
}
