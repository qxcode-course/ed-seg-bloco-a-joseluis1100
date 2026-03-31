package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func maxlenpath(matrix [][]int, posAtual Pos, valAnt int) int {
	if posAtual.x < 0 || posAtual.y < 0 || posAtual.x >= len(matrix) || posAtual.y >= len(matrix[0]) || matrix[posAtual.x][posAtual.y] <= valAnt  {
		return 0
	}
	cima := maxlenpath(matrix, Pos{posAtual.x-1, posAtual.y}, matrix[posAtual.x][posAtual.y])
	baixo := maxlenpath(matrix, Pos{posAtual.x+1, posAtual.y}, matrix[posAtual.x][posAtual.y])
	esquerda := maxlenpath(matrix, Pos{posAtual.x, posAtual.y-1}, matrix[posAtual.x][posAtual.y])
	direita := maxlenpath(matrix, Pos{posAtual.x, posAtual.y+1}, matrix[posAtual.x][posAtual.y])
	dir := cima
	if baixo > dir {
		dir = baixo
	}
	if esquerda > dir {
		dir = esquerda
	}
	if direita > dir {
		dir = direita
	}
	return 1 + dir
}

func longestIncreasingPath(matrix [][]int) int {
	maior := 1
	for i := range matrix {
		for j := range matrix[0] {
			caminho := maxlenpath(matrix, Pos{i, j}, -1)
			if caminho > maior {
				maior = caminho
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