package main

import (
	"fmt"
	"strconv"
	"strings"
)

func iniciarSoldados(N int) []bool {
	soldados := make([]bool, N)
	for i := range soldados {
		soldados[i] = true
	}
	return soldados
}

func imprimirSoldados(E int, soldados []bool) {
	vetor := make([]string, 0)
	for i := range soldados {
		if i == E {
			vetor = append(vetor, strconv.Itoa(i+1)+">")
		} else if soldados[i] {
			vetor = append(vetor, strconv.Itoa(i+1))
		}
	}
	fmt.Printf("[ %s ]\n", strings.Join(vetor, " "))
}

func proxVivo(pos int, soldados []bool) int {
	for i, j := (pos+1)%len(soldados), 0; j < len(soldados); i, j = (i+1)%len(soldados), j+1 {
		if soldados[(i)] {
			return i
		}
	}
	return -1
}

func jogo(E int, soldados []bool) {
	prox := proxVivo(E, soldados)
	if prox == -1 {
		return
	}
	imprimirSoldados(E, soldados)
	soldados[prox] = false
	jogo(proxVivo(E, soldados), soldados)
}

func main() {
	var N, E int
	fmt.Scan(&N, &E)
	soldados := iniciarSoldados(N)
	jogo(E-1, soldados)
}
