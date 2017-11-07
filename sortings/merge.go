package sortings

// Merge sorts slice of integers
func Merge(numbers []int) {
	sort(mainMerge, numbers)
}

func mainMerge(numbers []int) {
	mergeSort(numbers, 0, len(numbers)-1)
}

func merge(numbers []int, l, m, r int) {
	helper := make([]int, len(numbers))

	for i := l; i <= r; i++ {
		helper[i] = numbers[i]
	}

	helperLeft, helperRight, current := l, m+1, l

	for helperLeft <= m && helperRight <= r {
		if helper[helperLeft] <= helper[helperRight] {
			numbers[current] = helper[helperLeft]
			helperLeft++
		} else {
			numbers[current] = helper[helperRight]
			helperRight++
		}

		current++
	}

	remaining := m - helperLeft

	for i := 0; i <= remaining; i++ {
		numbers[current+i] = helper[helperLeft+i]
	}
}

func mergeSort(numbers []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(numbers, l, m)
		mergeSort(numbers, m+1, r)
		merge(numbers, l, m, r)
	}
}
