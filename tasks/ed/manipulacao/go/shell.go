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

func getMen(vet []int) []int {
	men := make([]int, 0)
	for _, valor := range vet {
		if valor > 0 {
			men = append(men, valor)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	calmWomen := make([]int, 0)
	for _, valor := range vet {
		if valor < 0 && valor > -10 {
			calmWomen = append(calmWomen, valor)
		}
	}
	return calmWomen

}

func sortVet(vet []int) []int {
	sorted := slices.Clone(vet)
	slices.Sort(sorted)
	return sorted
}

func sortStress(vet []int) []int {
	index := 1
	for index < len(vet) {
		if index == 0 || math.Abs(float64(vet[index])) >= math.Abs(float64(vet[index-1])) {
			index++
		} else {
			vet[index], vet[index-1] = vet[index-1], vet[index]
			index--
		}
	}
	return vet
}

func reverse(vet []int) []int {
	reversed := slices.Clone(vet)
	slices.Reverse(reversed)
	return reversed
}

func unique(vet []int) []int {
	unicos := make([]int, 0)
	mapa := make(map[int]int)
	for _, valor := range vet {
		mapa[valor]++
		if mapa[valor] == 1 {
			unicos = append(unicos, valor)
		}
	}
	return unicos
}

func repeated(vet []int) []int {
	repetidos := make([]int, 0)
	mapa := make(map[int]int)
	for _, valor := range vet {
		mapa[valor]++
		if mapa[valor] > 1 {
			repetidos = append(repetidos, valor)
		}
	}
	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
