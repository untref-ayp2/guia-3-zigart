package linkedlist

import "errors"

type Queue[T comparable] struct {
	items LinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
	queue := new(Queue[T])
	return queue
}

func (q *Queue[T]) Enqueue(v T) {
	q.items.Append(v)
}

func (q *Queue[T]) Dequeue() {
	if !q.IsEmpty() {
		// fmt.Print(q.items.head.value)
		q.items.Remove(q.items.head.value)
	}
}

func (q *Queue[T]) Front() (T, error) {
	if !q.IsEmpty() {
		return q.items.head.value, nil
	}
	var t T
	return t, errors.New("La lista esta vacia \n")
}

func (q *Queue[T]) IsEmpty() bool {
	return q.items.size == 0
}
