package linkedlist

import "errors"

type Stack[T comparable] struct {
	items LinkedList[T]
}

func NewStack[T comparable]() *Stack[T] {
	stack := new(Stack[T])
	return stack
}

func (q *Stack[T]) Push(v T) {
	q.items.Append(v)
}

func (q *Stack[T]) Pop() {
	if !q.IsEmpty() {
		q.items.Remove(q.items.tail.value)
	}
}

func (q *Stack[T]) Front() (T, error) {
	if !q.IsEmpty() {
		return q.items.head.value, nil
	}
	var t T
	return t, errors.New("La lista esta vacia \n")
}

func (q *Stack[T]) IsEmpty() bool {
	return q.items.size == 0
}
