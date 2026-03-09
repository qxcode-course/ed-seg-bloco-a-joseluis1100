package main

import "fmt"

func operacao(valor int) {
	if valor == 0 {
		return
	}
	divInt := valor / 2
	resto := valor % 2
	operacao(divInt)
	fmt.Printf("%d %d\n", divInt, resto)
}

func main() {
	var valor int
	fmt.Scan(&valor)
	operacao(valor)
}
