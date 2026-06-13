package main

import (
	"fmt"
	"strconv"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	curr := l.root.next
	str := "[ "
	for ; curr != l.root; curr = curr.next {
		str += strconv.Itoa(curr.Value)
		if curr == sword {
			str += ">"
		}
		str += " "
	}
	str += "]"
	return str
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	it = it.next
	if it == l.root {
		it = it.next
	}
	return it
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
