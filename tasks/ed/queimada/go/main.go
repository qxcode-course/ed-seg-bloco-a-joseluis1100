package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	//_, _, _ = mat, l, c
	// se estiver fora da matriz, retorne
	if l >= len(grid) || l < 0 || c >= len(grid[0]) || c < 0 {
		return
	}
	// se o elemento atual não for uma arvore, retorne
	if grid[l][c] != '#' {
		return
	}
	// queime a arvore colocando o caractere 'o' na posição atual
	grid[l][c] = 'o'
	// chame a recursão para todos os 4 vizinhos
	burnTrees(grid, l+1, c)
	burnTrees(grid, l-1, c)
	burnTrees(grid, l, c+1)
	burnTrees(grid, l, c-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
