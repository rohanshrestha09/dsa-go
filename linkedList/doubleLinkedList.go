package linkedList

import "fmt"

func doubleLinkedList() {
	type node struct {
		data int
		prev *node
		next *node
	}

	var root *node

	createNode := func() (node, bool) {
		var data int

		fmt.Print("Enter a data to insert: ")

		_, err := fmt.Scan(&data)

		if err != nil {
			panic(err)
		}

		newNode := node{data: data, prev: nil, next: nil}

		if root != nil {
			return newNode, false
		}

		root = &newNode

		return newNode, true
	}
}
