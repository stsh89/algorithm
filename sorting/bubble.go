package sorting

func Bubble(numbers []int) {
	if len(numbers) < 2 {
		return
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}

			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
}
