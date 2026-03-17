package main

import (
	"bufio"
	"fmt"
	"os"
)

func busca(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	busca(grid, x-1, y)
	busca(grid, x+1, y)
	busca(grid, x, y+1)
	busca(grid, x, y-1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	var qtd_ilhas int
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				qtd_ilhas++
				busca(grid, i, j)
			}
		}
	}
	return qtd_ilhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
