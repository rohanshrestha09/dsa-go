package queue

import "fmt"

type Queue[T any] struct {
	array []T
}

func (queue *Queue[T]) Enqueue(data T) {
	queue.array = append(queue.array, data)
}

func (queue *Queue[T]) Dequeue() T {
	if len(queue.array) == 0 {
		panic("Queue is empty")
	}

	queue.array = queue.array[1:]

	return queue.array[0]
}

func (queue *Queue[T]) Display() {
	for _, val := range queue.array {
		fmt.Print(val, " ")
	}
}
