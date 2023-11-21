package sorting

func BubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {

		swap := false

		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				temp := data[j]

				data[j] = data[j+1]

				data[j+1] = temp

				swap = true
			}
		}

		if !swap {
			break
		}
	}

	return data
}
