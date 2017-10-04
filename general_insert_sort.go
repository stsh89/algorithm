package algorithm

func GeneralInsertSort(array []int, order bool) []int {
	if len(array) < 2 {
		return array
	}

	for i := 1; i < len(array); i++ {
		j, tmp := i, array[i]

		for {
			if (j > 0) && (sortingCondition(array[j-1], tmp, order)) {
				array[j] = array[j-1]
				array[j-1] = tmp
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
