package unionfind_test

import (
	uf "github.com/rohanshrestha09/dsa-go/union-find"
	"testing"
)

func TestPercolation(t *testing.T) {
	p := new(uf.Percolation)

	p.Init(5)

	openSites := [][2]int{
		{1, 1}, {1, 2}, {1, 4},
		{2, 4},
		{3, 2}, {3, 4}, {3, 5},
		{4, 1}, {4, 3},
		{5, 1}, {5, 2}, {5, 4}, {5, 5},
	}

	for _, v := range openSites {
		p.Open(v[0], v[1])
	}

	/*
			   1 1 0 1 0
			   0 0 0 1 0
			   0 1 0 1 1
			   1 0 1 0 0
			   1 1 0 1 1

		    1 represents open site and 0 represents closed site
	*/

	t.Run("should not percolate", func(t *testing.T) {
		percolates := p.Percolates()

		if percolates {
			t.Errorf("should not percolate")
		}
	})

	t.Run("count open sites", func(t *testing.T) {
		expected := len(openSites)

		result := p.NumberOfOpenSites()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	p.Open(3, 3)

	p.Open(4, 2)

	/*
			   1 1 0 1 0
			   0 0 0 1 0
			   0 1 1 1 1
			   1 1 1 0 0
			   1 1 0 1 1

		    1 represents open site and 0 represents closed site
	*/

	t.Run("should percolate", func(t *testing.T) {
		percolates := p.Percolates()

		if !percolates {
			t.Errorf("should percolate")
		}
	})
}
