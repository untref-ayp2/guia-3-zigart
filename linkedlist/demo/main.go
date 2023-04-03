package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {
	l := linkedlist.NewLinkedList[int]()
	fmt.Println("Agregamos 1, 2 y 3 al final de la lista")
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l)
	fmt.Println("Agregamos 0 al inicio de la lista")
	l.Prepend(0)
	fmt.Println(l)
	fmt.Println("Buscamos el numero 3")
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))

	l.Remove(2)
	l.Remove(0)
	l.Remove(3)
	fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	fmt.Println(l)
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))
	l.Remove(1)
	l.Remove(1) //No deberia hacer nada
	l.Remove(1) //No deberia hacer nada
	fmt.Println(l)

	fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	l.Append(0)
	l.Append(1)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	fmt.Println(l)
	fmt.Println("Agregamos 2 en la posicion 2")
	l.InsertAt(2, 2)

	fmt.Println(l)

	l2 := linkedlist.NewLinkedList[int]()
	l2.Append(100)
	l2.Append(200)
	l2.Append(300)

	result, _ := l.ConcatenarLinkedList(l2)
	fmt.Print(result)

}

func ConcatenarLinkedList[T comparable](l1 *linkedlist.LinkedList[T], l2 *linkedlist.LinkedList[T]) *linkedlist.LinkedList[T] {

	aux := linkedlist.NewLinkedList[T]()

	for i := 0; i < l1.Size(); i++ {

		dato, _ := l1.Get(i)

		aux.Append(dato)
	}

	for i := 0; i < l2.Size(); i++ {
		dato, _ := l2.Get(i)

		aux.Append(dato)
	}

	return aux
}
