package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	vetS := make([]string, len(vet))
	for i, valor := range vet {
		vetS[i] = strconv.Itoa(valor)
	}
	return "[" + strings.Join(vetS, ", ") + "]"
}

func tostrrev(vet []int) string {
	reverse(vet)
	return tostr(vet)
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	slices.Reverse(vet)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var soma int
	for _, valor := range vet {
		soma += valor
	}
	return soma
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	mul := 1
	for _, valor := range vet {
		mul *= valor
	}
	return mul
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	var rec func(v []int) (int, int)
	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return v[0], 0
		}
		valor, i := rec(v[1:])
		i++
		if v[0] < valor {
			return v[0], 0
		}
		return valor, i
	}
	_, imenor := rec(vet)
	return imenor
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
