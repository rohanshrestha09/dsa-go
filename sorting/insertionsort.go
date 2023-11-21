package sorting

func InsertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		k := i

		for k-1 >= 0 && data[k] < data[k-1] {
			temp := data[k-1]

			data[k-1] = data[k]

			data[k] = temp

			k--
		}
	}

	return data
}
