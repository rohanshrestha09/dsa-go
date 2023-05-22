package sorting

import "fmt"

func insertionSort(data []Song) {

	fmt.Println("Before:", data)

	for i := 1; i < len(data); i++ {
		k := i

		for k-1 >= 0 && data[k].ReleaseYear > data[k-1].ReleaseYear {
			temp := data[k-1]

			data[k-1] = data[k]

			data[k] = temp

			k--
		}
	}

	fmt.Println("After:", data)
}
