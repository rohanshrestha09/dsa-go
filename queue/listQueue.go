package queue

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type ListQueue[T any] struct {
	front  *Node[T]
	rear   *Node[T]
	Length int
}

func (q *ListQueue[T]) Enqueue(data T) {
	newNode := &Node[T]{data, nil}

	if q.front == nil && q.rear == nil {
		q.front = newNode

		q.rear = newNode

		q.Length++

		return
	}

	q.rear.next = newNode

	q.rear = newNode

	q.Length++
}

func (q *ListQueue[T]) Dequeue() T {
	if q.front == nil && q.rear == nil {
		panic("Queue is empty")
	}

	temp := q.front

	q.front = q.front.next

	temp.next = nil

	if q.front == nil {
		q.rear = nil
	}

	q.Length--

	return temp.data
}

func (q *ListQueue[T]) Display() {
	temp := q.front

	for temp != nil {
		fmt.Printf("%v\t", temp.data)
	}
}
