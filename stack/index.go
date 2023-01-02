package stack

import "fmt"

func Run() {

	var (
		arr    []int
		choice int
	)

	push := func() {
		var data int

		fmt.Print("Enter a data to push: ")

		_, err := fmt.Scanf("%d", &data)
		if err != nil {
			panic("Something went wrong")
		}

		arr = append(arr, data)
	}

	pop := func() {
		fmt.Println("The popped element is ", arr[len(arr)-1])

		arr = arr[:len(arr)-1]
	}

	display := func() {
		for _, val := range arr {
			fmt.Print(val, " ")
		}
	}

	for true {

		fmt.Print("\n1.Push\n2.Pop\n3.Display")

		fmt.Print("\nEnter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			push()

		case 2:
			pop()

		case 3:
			display()
		}
	}

}
