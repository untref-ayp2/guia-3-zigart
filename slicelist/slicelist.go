package slicelist

import (
	"errors"
	"fmt"
)

type SliceList[T comparable] struct {
	slice []T
}

func NewSliceList[T comparable]() *SliceList[T] {
	return &SliceList[T]{}
}

func (l *SliceList[T]) Append(item T) {
	l.slice = append(l.slice, item)
}

func (l *SliceList[T]) Prepend(item T) {
	l.slice = append([]T{item}, l.slice...)
}

func (l *SliceList[T]) InsertAt(posicion int, item T) {
	l.slice = append(l.slice[:posicion], append([]T{item}, l.slice[posicion:]...)...)
}

func (l *SliceList[T]) Remove(item T) bool {
	for i := 0; i < len(l.slice); i++ {
		if l.slice[i] == item {
			l.slice = append(l.slice[:i], l.slice[i+1:]...)
			return true
		}
	}
	return false
}

func (l *SliceList[T]) Get(posicion int) (T, error) {
	if posicion < 0 || posicion >= len(l.slice) {
		var t T
		return t, errors.New("Lista vacía")
	}
	return l.slice[posicion], nil
}

func (l *SliceList[T]) Size() int {
	return len(l.slice)
}

func (l *SliceList[T]) String() string {
	if l.Size() == 0 {
		return "[]"
	}
	result := "["
	for i := 0; i < l.Size(); i++ {
		obtenido, err := l.Get(i)
		_, errorProximaPosicion := l.Get(i + 1)
		if err == nil && errorProximaPosicion == nil {
			result += fmt.Sprintf("%v,", obtenido)
		} else {
			result += fmt.Sprintf("%v", obtenido)
		}
	}
	result += "]"
	return result
}
