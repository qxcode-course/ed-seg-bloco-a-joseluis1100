package main

import (
	"fmt"
	"slices"
)

func scanear(n int) []int {
	vetor := make([]int, n)
	for i := range vetor {
		fmt.Scan(&vetor[i])
	}
	return vetor
}

func busca(numeros []int, k, i, sum int) bool {
	if sum == k {
		return true
	}
	if i == len(numeros) || sum > k {
		return false
	}
	return busca(numeros, k, i+1, sum + numeros[i]) || busca(numeros, k, i+1, sum)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	numeros := scanear(n)
	slices.Sort(numeros)
	fmt.Println(busca(numeros, k, 0, 0))
}
