package queue

type Queue []any

func (q *Queue) Enqueue(data any) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() any {
	item := (*q)[0]

	*q = (*q)[1:]

	return item
}
