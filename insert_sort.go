package algorithm

func InsertSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	for i := 1; i < len(array); i++ {
		j, tmp := i, array[i]

		for {
			if (j > 0) && (array[j-1] > tmp) {
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
