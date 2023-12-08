package sorting

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)

	result = append(result, right[j:]...)

	return result
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := MergeSort(data[:mid])

	right := MergeSort(data[mid:])

	return merge(left, right)
}
