package sorting

func Insert(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	for i := 1; i < len(numbers); i++ {
		j, tmp := i, numbers[i]

		for {
			if (j > 0) && (numbers[j-1] > tmp) {
				numbers[j] = numbers[j-1]
				numbers[j-1] = tmp
				j--
			} else {
				break
			}
		}
	}

	return numbers
}
