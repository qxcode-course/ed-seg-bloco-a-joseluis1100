package main

import "fmt"

func adicionarAnimal(animais map[int]int) int {
	var animal int
	fmt.Scan(&animal)
	if animais[-animal] > 0 {
		animais[-animal]--
		return 1
	} else {
		animais[animal]++
		return 0
	}
}

func main() {
	var n, casais int
	animais := make(map[int]int)
	fmt.Scan(&n)
	for range n {
		casais += adicionarAnimal(animais)
	}
	fmt.Println(casais)
}
