package queue

type Queue[T any] struct {
	items []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	return q.items[len(q.items)-1]
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Clear() {
	q.items = []T{}
}
