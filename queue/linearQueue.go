package queue

import "fmt"

func linearQueue() {
	const SIZE = 5

	var (
		QUEUE       [SIZE]int
		choice      int
		front, rear = -1, -1
	)

	enqueue := func() {
		var data int

		fmt.Print("Enter a data to enqueue: ")

		_, err := fmt.Scanf("%d", &data)

		if err != nil {
			panic(err)
		}

		switch {
		case rear == SIZE-1:
			fmt.Println("Queue full")

		case front == -1 && rear == -1:
			front++
			rear++
			QUEUE[rear] = data

		default:
			rear++
			QUEUE[rear] = data
		}
	}

	dequeue := func() {

		switch {
		case front == -1 && rear == -1:
			fmt.Println("Queue is empty")

		case front == rear:
			fmt.Println("The dequeued element is ", QUEUE[front])
			front = -1
			rear = -1

		default:
			fmt.Println("The dequeued element is ", QUEUE[front])
			front++
		}
	}

	display := func() {
		if front == -1 && rear == -1 {
			fmt.Println("Queue is empty")
			return
		}

		for i := front; i <= rear; i++ {
			fmt.Print(QUEUE[i], " ")
		}
	}

	for true {

		fmt.Print("\n1.Enqueue\n2.Dequeue\n3.Display")

		fmt.Print("\nEnter your choice: ")

		_, err := fmt.Scan(&choice)

		if err != nil {
			return
		}

		switch choice {
		case 1:
			enqueue()

		case 2:
			dequeue()

		case 3:
			display()
		}
	}
}
