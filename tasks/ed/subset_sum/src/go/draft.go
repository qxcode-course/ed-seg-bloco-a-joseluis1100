package main

import "fmt"

func scanear(n int) []int {
	vetor := make([]int, n)
	for i := range vetor {
		fmt.Scan(&vetor[i])
	}
	return vetor
}

func busca(numeros []int, k int) bool {
	for i := range numeros {
		for j := range numeros[i+1:] {
			if numeros[i]+numeros[j] == k {
				return true
			}
		}
	}
	return false
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	numeros := scanear(n)
	fmt.Println(busca(numeros, k))
}
