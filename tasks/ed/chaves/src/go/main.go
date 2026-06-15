package main

import "fmt"

func main() {
	jogos := make([]int, 15)
	for i := range jogos {
		var mm, mn int
		fmt.Scan(&mm, &mn)
		jogos[i] = mm - mn
	}
	jogadores := NewQueue[int]()
	for i := range 16 {
		jogadores.Enqueue(i+1)
	}
	for i := range 15 {
		jog1, jog2 := jogadores.Dequeue(), jogadores.Dequeue()
		if jogos[i] > 0 {
			jogadores.Enqueue(jog1)
		} else {
			jogadores.Enqueue(jog2)
		}
	}
	fmt.Printf("%c\n", jogadores.Dequeue() + 'A' - 1)
}
