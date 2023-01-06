package circularLinkedList

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) InsertAtStart(newNode *Node) {
	if list.head == nil {
		list.head, list.tail = newNode, newNode
		list.length++
		return
	}

	list.head.prev = newNode

	newNode.next = list.head

	newNode.prev = list.tail

	list.head = newNode

	list.length++
}

func (list *LinkedList) InsertAtEnd(newNode *Node) {
	if list.head == nil {
		list.head, list.tail = newNode, newNode
		list.length++
		return
	}

	newNode.prev = list.tail

	newNode.next = list.head

	list.tail.next = newNode

	list.tail = newNode

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

	temp := list.head

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
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	if list.head == list.tail {
		list.head, list.tail = nil, nil
		list.length--
		return
	}

	temp := list.head

	list.head = list.head.next

	list.head.prev = list.tail

	list.tail.next = (func() *Node {
		if list.head == list.tail {
			return nil
		}
		return list.head
	})()

	temp.next, temp.prev = nil, nil

	list.length--
}

func (list *LinkedList) DeleteAtEnd() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	if list.head == list.tail {
		list.head, list.tail = nil, nil
		list.length--
		return
	}

	temp := list.tail

	list.tail = list.tail.prev

	list.tail.next = (func() *Node {
		if list.head == list.tail {
			return nil
		}
		return list.head
	})()

	temp.prev, temp.next = nil, nil

	list.length--
}

func (list *LinkedList) DeleteAtSpecific(pos int) {
	if list.head == nil {
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

	temp := list.head

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
	temp := list.head

	for temp != nil {
		fmt.Printf("%d\t", temp.data)
		temp = temp.next

		if temp == list.head {
			break
		}
	}
}

func (list *LinkedList) Search(data int) {
	temp := list.head

	for temp != nil {
		if temp.data == data {
			fmt.Println("Data found")
			return
		}

		temp = temp.next

		if temp == list.head {
			break
		}
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
