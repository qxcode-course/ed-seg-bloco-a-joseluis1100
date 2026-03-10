package main

import (
	"bufio"
	"fmt"
	"os"
)

func verificaIgual(index int, vetor1, vetor2 string) string {
	if index == len(vetor1) {
		return "iguais"
	}
	if vetor1[index] != vetor2[index] || len(vetor1) != len(vetor2) {
		return "diferentes"
	}
	index++
	
	return verificaIgual(index, vetor1, vetor2)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vetor1 := scanner.Text()
	scanner.Scan()
	vetor2 := scanner.Text()
	fmt.Println(verificaIgual(0, vetor1, vetor2))
}
