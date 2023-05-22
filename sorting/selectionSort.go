package sorting

import "fmt"

func selectionSort(data []Song) {
	fmt.Println("Before:", data)

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].TimesPlayed < data[j].TimesPlayed {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}

	fmt.Println("After:", data)

}
