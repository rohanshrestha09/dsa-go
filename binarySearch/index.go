package binarysearch

import (
	"fmt"
)

type sortOrder string

const (
	asc  sortOrder = "ASCENDING"
	desc sortOrder = "DESCENDING"
)

func isSorted(array []int, order sortOrder) bool {

	if len(array) == 1 {
		return true
	}

	if order == asc {
		if array[0] < array[1] {
			return isSorted(array[1:], order)
		}
	} else {
		if array[0] > array[1] {
			return isSorted(array[1:], order)
		}
	}

	return false

}

func binarySearch(array []int, item int) bool {

	if len(array) == 1 {
		return array[0] == item
	}

	mid := len(array) / 2

	guess := array[mid]

	if guess == item {
		return true
	}

	if guess > item {
		return binarySearch(array[:mid], item)
	}

	return binarySearch(array[mid:], item)

}

func Run() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	hasElement := binarySearch(array, 100)

	fmt.Println(hasElement)

	fmt.Println(isSorted(array, asc))

}
