package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adicionarConsultados(numConsultados int) map[string]int {
	consultados := make(map[string]int)
	var consultado string
	for range numConsultados {
		fmt.Scan(&consultado)
		consultados[consultado]++
	}
	return consultados
}

func adicionarBuscas(numBuscas int) []string {
	buscas := make([]string, numBuscas)
	for i := range buscas {
		fmt.Scan(&buscas[i])
	}
	return buscas
}

func procurarMatchs(consultados map[string]int, buscas []string) string {
	resultados := make([]string, len(buscas))
	for i, busca := range buscas {
		resultados[i] = strconv.Itoa(consultados[busca])
	}
	return strings.Join(resultados, " ")
}

func main() {
	var numConsultados, numBuscas int
	fmt.Scan(&numConsultados)
	consultados := adicionarConsultados(numConsultados)
	fmt.Scan(&numBuscas)
	buscas := adicionarBuscas(numBuscas)
	fmt.Println(procurarMatchs(consultados, buscas))
}
