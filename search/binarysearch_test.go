package search_test

import (
	"testing"

	s "github.com/rohanshrestha09/dsa-go/search"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	searchElement := 8

	result := s.BinarySearch(data, searchElement)

	if !result {
		t.Errorf("%d must be present %v", searchElement, data)
	}
}
