package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	if grid[l][c] != '#' {
		return
	}
	stack := NewStack[Pos]()
	stack.Push(Pos{l: l, c: c})
	for !stack.IsEmpty() {
		tree := stack.Pop()
		grid[tree.l][tree.c] = 'o'
		if tree.l + 1 < len(grid) && grid[tree.l+1][tree.c] == '#' {
			stack.Push(Pos{l: tree.l+1, c: tree.c})
		}
		if tree.l - 1 > -1 && grid[tree.l-1][tree.c] == '#' {
			stack.Push(Pos{l: tree.l-1, c: tree.c})
		}
		if tree.c + 1 < len(grid[0]) && grid[tree.l][tree.c+1] == '#' {
			stack.Push(Pos{l: tree.l, c: tree.c+1})
		}
		if tree.c - 1 > -1 && grid[tree.l][tree.c-1] == '#' {
			stack.Push(Pos{l: tree.l, c: tree.c-1})
		}
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

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

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
