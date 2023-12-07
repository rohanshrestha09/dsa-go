package queue

import "math/rand"

type RandomizedQueue[T any] []T

func (rq *RandomizedQueue[T]) Size() int {
	return len(*rq)
}

func (rq *RandomizedQueue[T]) IsEmpty() bool {
	return len(*rq) == 0
}

func (rq *RandomizedQueue[T]) Enqueue(item T) {
	*rq = append(*rq, item)
}

func (rq *RandomizedQueue[T]) Dequeue() T {
	randomIndex := rand.Intn(rq.Size())

	item := (*rq)[randomIndex]

	(*rq)[randomIndex] = (*rq)[rq.Size()-1]

	*rq = (*rq)[:rq.Size()-1]

	return item
}

func (rq *RandomizedQueue[T]) Sample() T {
	randomIndex := rand.Intn(rq.Size())

	return (*rq)[randomIndex]
}
