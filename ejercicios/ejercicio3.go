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

	l2 := linkedlist.NewLinkedList[int]()
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)
	l2.Append(8)
	l2.Append(10)

	/* concatenar(*l1, *l2) */
	l3 := intercalar(*l1, *l2)
	fmt.Println(l3.String())
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
		aux, err1 := l1.Get(i)
		if err1 == nil {
			l3.Append(int(aux))
		}
		aux, err1 = l2.Get(i)
		if err1 == nil {
			l3.Append(int(aux))
		}
	}
	return
}
