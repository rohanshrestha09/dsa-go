package sorting

import "fmt"

type Song struct {
	Name        string
	TimesPlayed uint
	ReleaseYear uint
	Genre       string
}

var data = []Song{
	{Name: "Timrai Tira", TimesPlayed: 50, ReleaseYear: 2023, Genre: "Romantic"},
	{Name: "Blue Eyes", TimesPlayed: 1, ReleaseYear: 2020, Genre: "Romantic"},
	{Name: "Timi Sangai", TimesPlayed: 90, ReleaseYear: 2023, Genre: "Romantic"},
	{Name: "Mayalu", TimesPlayed: 10, ReleaseYear: 2016, Genre: "HipHop"},
}

func Run() {
	fmt.Println(quickSort([]int{1, 4, 5, 10, 100, 99, 4, 5, 2000}))

	fmt.Println(bubbleSort(data))

	fmt.Println(selectionSort(data))

	fmt.Println(insertionSort(data))
}
