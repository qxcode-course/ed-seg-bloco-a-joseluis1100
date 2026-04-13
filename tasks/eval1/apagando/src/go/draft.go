package main

import (
	"fmt"
	"strings"
)

func montaFila(n int) []string {
	fila := make([]string, n)
	for i := range n {
		fmt.Scan(&fila[i])
	}
	return fila
}

func montaRemovidos(m int) map[string]int {
	var aux string
	removido := make(map[string]int)
	for range m {
		fmt.Scan(&aux)
		removido[aux]++
	}
	return removido
}

func toString(fila []string, removido map[string]int) string {
	vetString := make([]string, 0)
	for _, valor := range fila {
		if removido[valor] == 0 {
			vetString = append(vetString, valor)
		}
	}
	return strings.Join(vetString, " ")
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fila := montaFila(n)
	fmt.Scan(&m)
	removido := montaRemovidos(m)
	fmt.Println(toString(fila, removido) + " ")
}
