package sorting

func insertionSort(data []Song) []Song {

	for i := 1; i < len(data); i++ {
		k := i

		for k-1 >= 0 && data[k].ReleaseYear > data[k-1].ReleaseYear {
			temp := data[k-1]

			data[k-1] = data[k]

			data[k] = temp

			k--
		}
	}

	return data
}
