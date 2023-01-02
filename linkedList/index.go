package linkedlist

import "fmt"

type Node struct {
	data int
	link *Node
}

func SingleLinkedList() {
	root := Node{}

	fmt.Println(root)

	fmt.Println("Single Linked List")
}
