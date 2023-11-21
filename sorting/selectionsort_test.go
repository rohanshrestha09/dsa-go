package sorting_test

import (
	"sort"
	"testing"

	s "github.com/rohanshrestha09/dsa-go/sorting"
)

func TestSelectionSort(t *testing.T) {
	data := []int{1, 4, 5, 10, 100, 99, 4, 5, 2000}

	result := s.SelectionSort(append(make([]int, 0), data...))

	if !sort.IntsAreSorted(result) {
		sort.Ints(data)
		t.Errorf("Sorted array should be %v, not %v", data, result)
	}
}
