package queue_test

import (
	"testing"

	q "github.com/rohanshrestha09/dsa-go/queue"
)

func TestRandomizedQueue(t *testing.T) {
	var rq q.RandomizedQueue[int]

	size := 5

	for i := 1; i <= size; i++ {
		rq.Enqueue(i)
	}

	t.Run("test size after enqueue", func(t *testing.T) {
		expected := size

		result := rq.Size()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test size after dequeue", func(t *testing.T) {
		rq.Dequeue()

		rq.Dequeue()

		expected := 3

		result := rq.Size()

		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test size after removing all elements", func(t *testing.T) {
		rq.Dequeue()
		rq.Dequeue()
		rq.Dequeue()

		isEmpty := rq.IsEmpty()

		if !isEmpty {
			t.Errorf("expected queue to be empty")
		}
	})

}
