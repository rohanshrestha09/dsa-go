package queue

import "fmt"

type Node struct {
	data int
	next *Node
}

type ListQueue struct {
	front  *Node
	rear   *Node
	Length int
}

func (q *ListQueue) Enqueue(data int) {
	newNode := &Node{data, nil}

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

func (q *ListQueue) Dequeue() int {
	if q.front == nil && q.rear == nil {
		fmt.Println("Queue is empty")
		return -1
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

func (q *ListQueue) Display() {
	temp := q.front

	for temp != nil {
		fmt.Printf("%d\t", temp.data)
	}
}

func listQueue() {
	q := new(Queue)

	q.Enqueue(10)

	q.Enqueue(20)

	q.Enqueue(500)

	q.Dequeue()

	q.Display()
}
