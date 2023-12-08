package sorting

func ShellSort(data []int) []int {
	n := len(data)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := data[i]
			j := i

			for j >= gap && data[j-gap] > temp {
				data[j] = data[j-gap]
				j -= gap
			}

			data[j] = temp
		}

		gap /= 2
	}

	return data
}
