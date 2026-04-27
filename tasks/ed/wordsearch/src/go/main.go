package main

import (
	"bufio"
	"fmt"
	"os"
)

func seeNext(grid [][]byte, word string, index, x, y int) bool {
	if index == len(word) {
		return true
	}
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 || grid[x][y] != word[index] {
		return false
	}
	temp := grid[x][y]
	grid[x][y] = ' '
	if seeNext(grid, word, index+1, x-1, y) || seeNext(grid, word, index+1, x+1, y) ||
	seeNext(grid, word, index+1, x, y-1) || seeNext(grid, word, index+1, x, y+1) {
		return true
	}
	grid[x][y] = temp
	return false
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := range grid {
		for j := range grid[0] {
			if seeNext(grid, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
