package sorting

import (
	"math/rand"
	"sort"
)

func BogoSort(data []int) []int {

	for !sort.IntsAreSorted(data) {
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
	}

	return data
}
