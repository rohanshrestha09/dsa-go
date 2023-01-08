package queue

import "fmt"

type Queue struct {
	arr []int
}

func (queue *Queue) Enqueue(data int) {
	queue.arr = append(queue.arr, data)
}

func (queue *Queue) Dequeue() {
	if len(queue.arr) == 0 {
		fmt.Println("\nQueue is empty")
		return
	}

	fmt.Println("\nThe dequeued element is", queue.arr[0])

	queue.arr = queue.arr[1:]
}

func (queue *Queue) Display() {
	for _, val := range queue.arr {
		fmt.Print(val, " ")
	}
}

func circularQueue() {
	queue := &Queue{}

	queue.Enqueue(10)

	queue.Enqueue(20)

	queue.Enqueue(50)

	queue.Display() // 10 20 50

	queue.Dequeue() // The dequeued element is 10

	queue.Display() // 20 50
}
