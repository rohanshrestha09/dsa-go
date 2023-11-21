package sorting

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	var leftarray, rightarray []int

	mid := len(array) / 2

	pivot := array[mid]

	array = append(array[:mid], array[mid+1:]...)

	for _, val := range array {
		if val > pivot {
			rightarray = append(rightarray, val)
		} else {
			leftarray = append(leftarray, val)
		}
	}

	result := []int{}

	result = append(result, QuickSort(leftarray)...)

	result = append(result, pivot)

	result = append(result, QuickSort(rightarray)...)

	return result
}
