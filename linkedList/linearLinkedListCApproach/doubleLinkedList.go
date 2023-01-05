package linearLinkedListCApproach

import "fmt"

func doubleLinkedList() {
	type node struct {
		data int
		prev *node
		next *node
	}

	var root *node

	length := func() int {
		var count int

		temp := root

		for temp != nil {
			count++
			temp = temp.next
		}

		return count
	}

	createNode := func() (*node, bool) {
		var data int

		fmt.Print("Enter a data to insert: ")

		_, err := fmt.Scan(&data)

		if err != nil {
			panic(err)
		}

		newNode := &node{data: data, prev: nil, next: nil}

		if root != nil {
			return newNode, false
		}

		root = newNode

		return newNode, true
	}

	insertAtStart := func() {
		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		root.prev = newNode

		newNode.next = root

		root = newNode
	}

	insertAtEnd := func() {
		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		temp := root

		for temp.next != nil {
			temp = temp.next
		}

		newNode.prev = temp

		newNode.prev.next = newNode
	}

	insertAtSpecific := func() {
		var pos int

		fmt.Print("Enter the position to append: ")

		_, err := fmt.Scan(&pos)

		if err != nil {
			panic(err)
		}

		if pos == 1 {
			insertAtStart()
			return
		}

		if pos == length()+1 {
			insertAtEnd()
			return
		}

		if pos > length()+1 {
			fmt.Println("Position exceeds the list length")
			return
		}

		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		temp := root

		for i := 1; i < pos; i++ {
			temp = temp.next
		}

		temp.prev.next = newNode

		newNode.prev = temp.prev

		newNode.next = temp

		temp.prev = newNode

	}

	deleteAtStart := func() {
		if root == nil {
			fmt.Println("List is empty")
			return
		}

		temp := root

		root = root.next

		root.prev = nil

		temp.next = nil

		temp.prev = nil
	}

	deleteAtEnd := func() {
		if root == nil {
			fmt.Println("List is empty")
			return
		}

		if root.next == nil {
			root = nil
			return
		}

		temp := root

		for temp.next != nil {
			temp = temp.next
		}

		temp.prev.next = nil

		temp.prev = nil
	}

	deleteAtSpecific := func() {
		var pos int

		if root == nil {
			fmt.Println("List is empty")
			return
		}

		fmt.Print("Enter the position to delete: ")

		_, err := fmt.Scan(&pos)

		if err != nil {
			panic(err)
		}

		if pos == 1 {
			deleteAtStart()
			return
		}

		if length() == pos {
			deleteAtEnd()
			return
		}

		if pos > length() {
			fmt.Println("Position exceeds the list length")
			return
		}

		temp := root

		for i := 1; i < pos; i++ {
			temp = temp.next
		}

		temp.prev.next = temp.next

		temp.next.prev = temp.prev

		temp.next = nil

		temp.prev = nil
	}

	display := func() {
		temp := root

		for temp != nil {
			fmt.Printf("%d\t", temp.data)

			temp = temp.next
		}
	}

	search := func() {

		var data int

		fmt.Print("Enter data to search: ")

		_, err := fmt.Scan(&data)

		if err != nil {
			panic(err)
		}

		temp := root

		for temp != nil {
			if temp.data == data {
				fmt.Println("Data found")
				return
			}

			temp = temp.next
		}

		fmt.Println("Data not found")
	}

	for true {
		var choice int

		fmt.Print("\n1.Insert at start\n2.Insert at specific\n3.Insert at end\n4.Display\n5.Delete at start\n6.Delete at specific\n7.Delete at end\n8.Search\n9.Length of the list")

		fmt.Print("\nEnter your choice: ")

		_, err := fmt.Scan(&choice)

		if err != nil {
			panic(err)
		}

		switch choice {
		case 1:
			insertAtStart()
		case 2:
			insertAtSpecific()
		case 3:
			insertAtEnd()
		case 4:
			display()
		case 5:
			deleteAtStart()
		case 6:
			deleteAtSpecific()
		case 7:
			deleteAtEnd()
		case 8:
			search()
		case 9:
			fmt.Println("Length of the list is", length())
		}
	}
}
