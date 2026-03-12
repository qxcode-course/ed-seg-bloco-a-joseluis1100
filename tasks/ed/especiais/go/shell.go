package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	mapa := make(map[int]int)
	ocorrencias := make([]Pair, 0)
	for _, valor := range vet {
		mapa[int(math.Abs(float64(valor)))]++
	}
	index := 0
	for range slices.Max(vet) + 1 {
		if mapa[index] != 0 {
			ocorrencias = append(ocorrencias, Pair{index, mapa[index]})
		}
		index++
	}
	return ocorrencias
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}
	times := make([]Pair, 0)
	id, contador := vet[0], 0
	for _, valor := range vet {
		if valor == id {
			contador++
		} else {
			times = append(times, Pair{id, contador})
			id = valor
			contador = 1
		}
	}
	times = append(times, Pair{id, contador})
	return times
}

func mnext(vet []int) []int {
	if len(vet) == 1 {
		return []int{0}
	}
	novoVetor := make([]int, 0)
	for i := range vet {
		var esquerda, direita bool
		if i == 0 {
			if vet[1] < 0 {
				direita = true
			}
		} else if i == len(vet)-1 {
			if vet[len(vet)-2] < 0 {
				esquerda = true
			}
		} else {
			if vet[i-1] < 0 {
				esquerda = true
			}
			if vet[i+1] < 0 {
				direita = true
			}
		}
		if vet[i] > 0 && (esquerda || direita) {
			novoVetor = append(novoVetor, 1)
		} else {
			novoVetor = append(novoVetor, 0)
		}
	}
	return novoVetor
}

func alone(vet []int) []int {
	if len(vet) == 1 {
		return []int{1}
	}
	novoVetor := make([]int, 0)
	for i := range vet {
		var esquerda, direita bool
		if i == 0 {
			if vet[1] < 0 {
				direita = true
			}
		} else if i == len(vet)-1 {
			if vet[len(vet)-2] < 0 {
				esquerda = true
			}
		} else {
			if vet[i-1] < 0 {
				esquerda = true
			}
			if vet[i+1] < 0 {
				direita = true
			}
		}
		if vet[i] > 0 && !(esquerda || direita) {
			novoVetor = append(novoVetor, 1)
		} else {
			novoVetor = append(novoVetor, 0)
		}
	}
	return novoVetor
}

func couple(vet []int) int {
	var casais int
	mapa := make(map[int]int)
	for _, valor := range vet {
		if mapa[-valor] > 0 {
			mapa[-valor]--
			casais++
		} else {
			mapa[valor]++
		}
	}
	return casais
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	for i := range seq {
		if pos+i >= len(vet) || !(vet[pos+i] == seq[i]) {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	for i := range vet {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	novoVetor := make([]int, 0)
	mapa := make(map[int]int)
	for _, valor := range posList {
		mapa[valor]++
	}
	for i := range vet {
		if mapa[i] == 0 {
			novoVetor = append(novoVetor, vet[i])
		}
	}
	return novoVetor
}

func clear(vet []int, value int) []int {
	novoVetor := make([]int, 0)
	for _, valor := range vet {
		if valor != value {
			novoVetor = append(novoVetor, valor)
		}
	}
	return novoVetor
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
