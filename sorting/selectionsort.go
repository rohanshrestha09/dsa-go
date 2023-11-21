package sorting

func SelectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}

	return data
}
