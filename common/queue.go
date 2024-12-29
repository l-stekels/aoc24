package common

type Queue[T interface{}] struct {
	items []T
	size  int
}

func NewQueue[T interface{}]() Queue[T] {
	return Queue[T]{items: []T{}}
}

func (q *Queue[T]) Enqueue(entry T) {
	q.items = append([]T{entry}, q.items...)
	q.size++
}

func (q *Queue[T]) Dequeue() T {
	if q.size == 0 {
		panic("queue is empty")
	}
	entry := q.items[0]
	q.items = q.items[1:]
	q.size--

	return entry
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
