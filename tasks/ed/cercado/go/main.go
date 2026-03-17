package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

func buscaBordas(mapa map[Pos]bool, board [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] == 'X' || mapa[Pos{x, y}] {
		return
	}
	mapa[Pos{x, y}] = true
	buscaBordas(mapa, board, x-1, y)
	buscaBordas(mapa, board, x+1, y)
	buscaBordas(mapa, board, x, y+1)
	buscaBordas(mapa, board, x, y-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visitados := make(map[Pos]bool)
	for i := range board {
		buscaBordas(visitados, board, i, 0)
		buscaBordas(visitados, board, i, len(board[0])-1)
	}
	for i := range board[0] {
		buscaBordas(visitados, board, 0, i)
		buscaBordas(visitados, board, len(board)-1, i)
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' && !visitados[Pos{i, j}] {
				board[i][j] = 'X'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
