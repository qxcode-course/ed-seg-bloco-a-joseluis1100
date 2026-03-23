package main

import (
	"bufio"
	"fmt"
	"os"
)

func buscaBarco(board [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] == '.' {
		return
	}
	board[x][y] = '.'
	buscaBarco(board, x-1, y)
	buscaBarco(board, x+1, y)
	buscaBarco(board, x, y-1)
	buscaBarco(board, x, y+1)
}	

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	contador := 0
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'X' {
				contador++
				buscaBarco(board, i, j)
			}
		}
	}
	return contador
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
