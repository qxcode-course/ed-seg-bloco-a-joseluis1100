package main

import "fmt"

func imprimir(s string, i int) {
	if i == -1 {
		return
	}
	fmt.Println(s[i:])
	imprimir(s, i-1)
}

func main() {
	var palavra string
	fmt.Scan(&palavra)
	imprimir(palavra, len(palavra)-1)
}
