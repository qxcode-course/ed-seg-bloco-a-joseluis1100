package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maiorPos(matrix [][]int, x, y, valAnt, cont int) (bool, int) {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || (valAnt >= matrix[x][y] && x != 0 && y != 0) {
		fmt.Print("vish ")
		return false, cont
	}
	fmt.Print("boa ")
	cont++
	esquerda, _ := maiorPos(matrix, x-1, y, matrix[x][y], cont)
	direita, _ := maiorPos(matrix, x+1, y, matrix[x][y], cont)
	cima, _ := maiorPos(matrix, x, y-1, matrix[x][y], cont)
	baixo, _ := maiorPos(matrix, x, y+1, matrix[x][y], cont)
	if esquerda || direita || cima || baixo {
		return true, cont
	}
	return false, cont
}

func longestIncreasingPath(matrix [][]int) int {
	maior := 0
	for i := range matrix {
		for j := range matrix[0] {
			_, maiorIndex := maiorPos(matrix, i, j, matrix[i][j], 0)
			if maiorIndex > maior {
				maior = maiorIndex
			}
		}
	}
	return maior
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
