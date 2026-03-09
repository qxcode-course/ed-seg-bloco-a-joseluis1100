package main

import (
	"fmt"
	"strings"
)

func montarFila(quantidadePessoas int) []string {
	pessoas := make([]string, quantidadePessoas)
	for i := range quantidadePessoas {
		fmt.Scan(&pessoas[i])
	}
	return pessoas
}

func mapearRemovidos(quantidadeRemovidos int) map[string]bool {
	var removido string
	removidos := make(map[string]bool)
	for range quantidadeRemovidos {
		fmt.Scan(&removido)
		removidos[removido] = true
	}
	return removidos
}

func novaFila(pessoas []string, removidos map[string]bool) string {
	fila := make([]string, 0)
	for _, pessoa := range pessoas {
		if !removidos[pessoa] {
			fila = append(fila, pessoa)
		}
	}
	return strings.Join(fila, " ")
}

func main() {
	var quantidadePessoas, quantidadeRemovidos int
	fmt.Scan(&quantidadePessoas)
	pessoas := montarFila(quantidadePessoas)
	fmt.Scan(&quantidadeRemovidos)
	removidos := mapearRemovidos(quantidadeRemovidos)
	fmt.Println(novaFila(pessoas, removidos) + " ")
}
