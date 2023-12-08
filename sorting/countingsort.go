package sorting

import (
	"slices"
)

func CountingSort(data []int) []int {
	max := slices.Max(data)

	count := make(map[int]int)

	result := make([]int, len(data))

	for _, v := range data {
		count[v] += 1
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	for i := len(data) - 1; i >= 0; i-- {
		result[count[data[i]]-1] = data[i]
		count[data[i]]--
	}

	return result
}
