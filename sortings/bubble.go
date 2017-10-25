package sortings

func Bubble(numbers []int) {
	sort(bubble, numbers)
}

func bubble(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		for j := len(numbers) - 1; j >= i; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}
}
