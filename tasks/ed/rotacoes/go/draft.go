package main

import (
	"fmt"
	"strconv"
	"strings"
)

func criarVetor(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range vetor {
		fmt.Scan(&vetor[i])
	}
	return vetor
}

func rotacionar(rotacoes int, vetor []int) string {
	vetorRotacionado := make([]string, len(vetor))
	for i := range vetor {
		vetorRotacionado[i] = strconv.Itoa(vetor[((i-rotacoes)%len(vetor)+len(vetor))%len(vetor)])
	}
	return strings.Join(vetorRotacionado, " ")
}

func main() {
	var tamanho, rotacoes int
	fmt.Scan(&tamanho, &rotacoes)
	vetor := criarVetor(tamanho)
	fmt.Printf("[ %s ]\n", rotacionar(rotacoes, vetor))
}
