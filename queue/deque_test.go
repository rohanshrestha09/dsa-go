package queue_test

import (
	"slices"
	"testing"

	q "github.com/rohanshrestha09/dsa-go/queue"
)

func TestDeque(t *testing.T) {
	dq := new(q.Deque[int])

	dq.AddFirst(1)

	dq.AddLast(2)

	dq.AddFirst(3)

	t.Run("check ordered elements", func(t *testing.T) {
		expected := []int{3, 1, 2}

		result := dq.ToSlice()

		if !slices.Equal(expected, result) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("remove first element", func(t *testing.T) {
		expected := 3

		result := dq.RemoveFirst()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("remove last element", func(t *testing.T) {
		expected := 2

		result := dq.RemoveLast()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("check if empty", func(t *testing.T) {
		isEmpty := dq.IsEmpty()

		if isEmpty {
			t.Errorf("should not be empty")
		}
	})

	t.Run("check size", func(t *testing.T) {
		expected := 1

		result := dq.Size()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}
