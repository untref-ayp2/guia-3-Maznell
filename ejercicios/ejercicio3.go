// Escribir una función que reciba como parámetros dos listas del mismo tipo y devuelva una
// lista resultante de concatenar la segunda lista al final de la primerapackage main
package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {
	l1 := linkedlist.NewLinkedList[int]()
	l1.Append(1)
	l1.Append(3)
	l1.Append(5)
	l1.Append(7)
	l1.Append(9)

	l2 := linkedlist.NewLinkedList[int]()
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)

	/* concatenar(*l1, *l2) */
	intercalar(*l1, *l2)

}

func concatenar(l1, l2 linkedlist.LinkedList[int]) (l3 linkedlist.LinkedList[int]) {

	l3 = l1
	for i := 0; i < l2.Size(); i++ {
		aux, _ := l2.Get(i)
		l3.Append(aux)
	}

	return

}

func intercalar(l1, l2 linkedlist.LinkedList[int]) (l3 linkedlist.LinkedList[int]) {
	count := 0
	if l1.Size() > l2.Size() {
		count = l1.Size()
	} else {
		count = l2.Size()
	}
	for i := 0; i < count; i++ {
		aux, _ := l1.Get(i)
		l3.Append(int(aux))
		aux, _ = l2.Get(i)
		l3.Append(int(aux))
	}
	fmt.Println(l3.String())
	return
}
