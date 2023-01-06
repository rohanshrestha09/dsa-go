//go:build exclude

package circularLinkedList

import "fmt"

type Node struct {
	data int
	link *Node
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

	newNode.link = list.head

	list.head = newNode

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

	temp = list.head

	for i := 1; i < pos; i++ {
		prev = temp

		temp = temp.link
	}

	newNode.link = temp

	prev.link = newNode

	list.length++
}

func (list *LinkedList) InsertAtEnd(newNode *Node) {
	if list.head == nil {
		list.InsertAtStart(newNode)
		return
	}

	list.tail.link = newNode

	newNode.link = list.head

	list.tail = newNode

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

	list.head = list.head.link

	list.tail.link = (func() *Node {
		if list.head == list.tail {
			return nil
		}
		return list.head
	})()

	temp.link = nil

	list.length--
}

func (list *LinkedList) DeleteAtSpecific(pos int) {

	switch {
	case list.head == nil:
		fmt.Println("List is empty")

	case pos == 1:
		list.DeleteAtStart()

	case list.length == pos:
		list.DeleteAtEnd()

	case pos > list.length:
		fmt.Println("Position exceeds the list length")

	default:
		var prev *Node

		temp := list.head

		for i := 1; i < pos; i++ {
			prev = temp
			temp = temp.link
		}

		prev.link = (func() *Node {
			if temp.link == list.head {
				return nil
			}
			return temp.link
		})()

		temp.link = nil

		list.length--
	}
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

	temp := list.head

	for temp.link != list.tail {
		temp = temp.link
	}

	list.tail = temp

	temp.link = (func() *Node {
		temp.link.link = nil
		if list.head == list.tail {
			return nil
		}
		return list.head
	})()

	list.length--
}

func (list *LinkedList) Display() {
	temp := list.head

	for temp != nil {
		fmt.Printf("%d\t", temp.data)
		temp = temp.link

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

		temp = temp.link

		if temp == list.head {
			break
		}
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
