package algorithm

func GeneralInsertSort(array *[]int, order bool) *[]int {
	arr := *array

	if len(arr) < 2 {
		return array
	}

	for i := 1; i < len(arr); i++ {
		j, tmp := i, arr[i]

		for {
			if (j > 0) && (sortingCondition(arr[j-1], tmp, order)) {
				arr[j] = arr[j-1]
				arr[j-1] = tmp
				j--
			} else {
				break
			}
		}
	}

	return array
}

func sortingCondition(prev, next int, order bool) bool {
	if order {
		return prev > next
	} else {
		return prev < next
	}

}
