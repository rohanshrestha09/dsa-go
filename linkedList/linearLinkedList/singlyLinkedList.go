//go:build exclude

package linearLinkedList

import "fmt"

type Node struct {
	data int
	link *Node
}

type LinkedList struct {
	root   *Node
	length int
}

func (list *LinkedList) InsertAtStart(newNode *Node) {
	if list.root != nil {
		newNode.link = list.root
	}

	list.root = newNode

	list.length++
}

func (list *LinkedList) InsertAtSpecific(newNode *Node, pos int) {
	if pos == 1 {
		list.InsertAtStart(newNode)
		return
	}

	if pos == list.length+1 {
		list.InsertAtEnd(newNode)
		return
	}

	if pos > list.length+1 {
		fmt.Println("Position exceeds the list length")

		return
	}

	var temp, prev *Node

	temp = list.root

	for i := 1; i < pos; i++ {
		prev = temp

		temp = temp.link
	}

	newNode.link = temp

	prev.link = newNode

	list.length++
}

func (list *LinkedList) InsertAtEnd(newNode *Node) {
	if list.root == nil {
		list.root = newNode
		list.length++
		return
	}

	temp := list.root

	for temp.link != nil {
		temp = temp.link
	}

	temp.link = newNode

	list.length++
}

func (list *LinkedList) DeleteAtStart() {
	if list.root == nil {
		fmt.Println("List is empty")

		return
	}

	temp := list.root

	list.root = list.root.link

	temp.link = nil

	list.length--
}

func (list *LinkedList) DeleteAtSpecific(pos int) {
	if list.root == nil {
		fmt.Println("List is empty")
		return
	}

	if pos == 1 {
		list.DeleteAtStart()
		return
	}

	if list.length == pos {
		list.DeleteAtEnd()
		return
	}

	if pos > list.length {
		fmt.Println("Position exceeds the list length")
		return
	}

	var prev *Node

	temp := list.root

	for i := 1; i < pos; i++ {
		prev = temp
		temp = temp.link
	}

	prev.link = temp.link

	temp.link = nil

	list.length--

}

func (list *LinkedList) DeleteAtEnd() {
	if list.root == nil {
		fmt.Println("List is empty")

		return
	}

	var prev *Node

	temp := list.root

	for temp.link != nil {
		prev = temp
		temp = temp.link
	}

	prev.link = nil

	list.length--
}

func (list *LinkedList) Display() {
	temp := list.root

	for temp != nil {
		fmt.Printf("%d\t", temp.data)
		temp = temp.link
	}
}

func (list *LinkedList) Search(data int) {
	temp := list.root

	for temp != nil {
		if temp.data == data {
			fmt.Println("Data found")
			return
		}

		temp = temp.link
	}

	fmt.Println("Data not found")
}

func singlyLinkedList() {
	linkedList := LinkedList{}

	linkedList.InsertAtStart(&Node{8, nil}) // 8

	linkedList.InsertAtStart(&Node{10, nil}) // 10 8

	linkedList.InsertAtEnd(&Node{20, nil}) // 10 8 20

	linkedList.DeleteAtEnd() // 10 8

	linkedList.InsertAtSpecific(&Node{30, nil}, 2) // 10 30 8

	linkedList.InsertAtSpecific(&Node{50, nil}, 3) // 10 30 50 8

	linkedList.DeleteAtSpecific(2) // 10 50 8

	linkedList.DeleteAtStart() // 50 8

	linkedList.Display()

	linkedList.Search(8) // Data found

	fmt.Println("\nlist of the length is", linkedList.length) // 2
}
