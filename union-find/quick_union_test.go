package unionfind_test

import (
	"testing"

	qf "github.com/rohanshrestha09/dsa-go/union-find"
)

func TestQuickUnionUF(t *testing.T) {
	qu := new(qf.QuickUnionUF)

	qu.Init(10)

	list := [][2]int{{0, 5}, {5, 6}, {1, 2}, {2, 7}, {3, 8}, {3, 4}, {4, 9}}

	for _, v := range list {
		qu.Union(v[0], v[1])
	}

	t.Run("test count component", func(t *testing.T) {
		expected := 3

		result := qu.Count()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	qu.Union(1, 6)

	t.Run("test count component after union of 2 component", func(t *testing.T) {
		expected := 2

		result := qu.Count()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test connection", func(t *testing.T) {
		p, q := 5, 6

		isConnected := qu.Connected(p, q)

		if !isConnected {
			t.Errorf("%d and %d should be connected", p, q)
		}
	})
}
