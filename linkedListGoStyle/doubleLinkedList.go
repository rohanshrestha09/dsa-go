package linkedListGoStyle

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	root   *Node
	length int
}

func (list *LinkedList) InsertAtStart(newNode *Node) {
	if list.root == nil {
		list.root = newNode
		list.length++
		return
	}

	list.root.prev = newNode

	newNode.next = list.root

	list.root = newNode

	list.length++
}

func (list *LinkedList) InsertAtEnd(newNode *Node) {
	if list.root == nil {
		list.root = newNode
		list.length++
		return
	}

	temp := list.root

	for temp.next != nil {
		temp = temp.next
	}

	newNode.prev = temp

	temp.next = newNode

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

	temp := list.root

	for i := 1; i < pos; i++ {
		temp = temp.next
	}

	temp.prev.next = newNode

	newNode.prev = temp.prev

	newNode.next = temp

	temp.prev = newNode

	list.length++
}

func (list *LinkedList) DeleteAtStart() {
	if list.root == nil {
		fmt.Println("List is empty")
		return
	}

	temp := list.root

	list.root = list.root.next

	list.root.prev = nil

	temp.next = nil

	list.length--
}

func (list *LinkedList) DeleteAtEnd() {
	if list.root == nil {
		fmt.Println("List is empty")
		return
	}

	if list.root.next == nil {
		list.root = nil
		list.length--
		return
	}

	temp := list.root

	for temp.next != nil {
		temp = temp.next
	}

	temp.prev.next = nil

	temp.prev = nil

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

	temp := list.root

	for i := 1; i < pos; i++ {
		temp = temp.next
	}

	temp.prev.next = temp.next

	temp.next.prev = temp.prev

	temp.next = nil

	temp.prev = nil

	list.length--
}

func (list *LinkedList) Display() {
	temp := list.root

	for temp != nil {
		fmt.Printf("%d\t", temp.data)

		temp = temp.next
	}
}

func (list *LinkedList) Search(data int) {
	if list.root == nil {
		fmt.Println("List is empty")
		return
	}

	temp := list.root

	for temp != nil {
		if temp.data == data {
			fmt.Println("Data found")
			return
		}

		temp = temp.next
	}

	fmt.Println("Data not found")

}

func doubleLinkedList() {

	linkedList := &LinkedList{}

	linkedList.InsertAtStart(&Node{data: 6}) // 6

	linkedList.InsertAtStart(&Node{data: 10}) // 10 6

	linkedList.InsertAtEnd(&Node{data: 14}) // 10 6 14

	linkedList.InsertAtSpecific(&Node{data: 20}, 2) // 10 20 6 14

	linkedList.InsertAtEnd(&Node{data: 50}) // 10 20 6 14 50

	linkedList.DeleteAtEnd() // 10 20 6 14

	linkedList.DeleteAtStart() // 20 6 14

	linkedList.DeleteAtSpecific(2) // 20 14

	linkedList.Display()

	linkedList.Search(14) // Data found

	fmt.Println("The length of the list is", linkedList.length) // 2
}
