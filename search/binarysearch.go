package search

type SortOrder int

const (
	ASC  SortOrder = 1
	DESC SortOrder = -1
)

func isSorted(data []int, order SortOrder) bool {
	if len(data) == 1 {
		return true
	}

	if order == ASC {
		if data[0] < data[1] {
			return isSorted(data[1:], order)
		}
	} else {
		if data[0] > data[1] {
			return isSorted(data[1:], order)
		}
	}

	return false
}

func BinarySearch(data []int, item int) bool {
	if len(data) == 1 {
		return data[0] == item
	}

	mid := len(data) / 2

	guess := data[mid]

	if guess == item {
		return true
	}

	if guess > item {
		return BinarySearch(data[:mid], item)
	}

	return BinarySearch(data[mid:], item)
}
