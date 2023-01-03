package queue

import "fmt"

func Run() {
	var (
		arr    []int
		choice int
	)

	enqueue := func() {
		var data int

		fmt.Print("Enter a data to enqueue: ")

		_, err := fmt.Scanf("%d", &data)
		if err != nil {
			panic(err)
		}

		arr = append(arr, data)
	}

	dequeue := func() {

		if len(arr) == 0 {
			fmt.Println("Queue is empty")
			return
		}

		fmt.Println("The dequeued element is ", arr[0])

		arr = arr[1:]
	}

	display := func() {
		for _, val := range arr {
			fmt.Print(val, " ")
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
