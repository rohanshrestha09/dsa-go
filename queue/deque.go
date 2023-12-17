package queue

type Node[T any] struct {
	item T
	prev *Node[T]
	next *Node[T]
}

type Deque[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (dq *Deque[T]) Enqueue(item T) {
	dq.AddLast(item)
}

func (dq *Deque[T]) Dequeue() T {
	return dq.RemoveFirst()
}

func (dq *Deque[T]) IsEmpty() bool {
	return dq.size == 0
}

func (dq *Deque[T]) Size() int {
	return dq.size
}

func (dq *Deque[T]) AddFirst(item T) {
	newNode := &Node[T]{item, nil, nil}

	if dq.head == nil {
		dq.head = newNode
		dq.tail = newNode
		dq.size++
		return
	}

	dq.head.prev = newNode

	newNode.next = dq.head

	newNode.prev = dq.tail

	dq.head = newNode

	dq.tail.next = newNode

	dq.size++
}

func (dq *Deque[T]) AddLast(item T) {
	newNode := &Node[T]{item, nil, nil}

	if dq.head == nil {
		dq.head = newNode
		dq.tail = newNode
		dq.size++
		return
	}

	dq.tail.next = newNode

	newNode.next = dq.head

	newNode.prev = dq.tail

	dq.tail = newNode

	dq.head.prev = newNode

	dq.size++
}

func (dq *Deque[T]) RemoveFirst() T {
	if dq.size == 0 {
		panic("No such element")
	}

	if dq.head == dq.tail {
		item := dq.head.item
		dq.head = nil
		dq.tail = nil
		dq.size--
		return item
	}

	temp := dq.head

	dq.head = dq.head.next

	dq.head.prev = dq.tail

	temp.next, temp.prev = nil, nil

	dq.size--

	return temp.item
}

func (dq *Deque[T]) RemoveLast() T {
	if dq.size == 0 {
		panic("No such element")
	}

	if dq.head == dq.tail {
		item := dq.head.item
		dq.head = nil
		dq.tail = nil
		dq.size--
		return item
	}

	temp := dq.tail

	dq.tail = dq.tail.prev

	temp.prev, temp.next = nil, nil

	dq.size--

	return temp.item
}

func (dq *Deque[T]) ToSlice() []T {

	items := []T{}

	temp := dq.head

	for temp != nil {
		items = append(items, temp.item)

		temp = temp.next

		if temp == dq.head {
			break
		}
	}

	return items
}
