package main

import "fmt"

type Path struct {
	x, y, dia int
}

func maximumMinutes(grid [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pessoa := make([][]int, len(grid))
	fogo := make([][]int, len(grid))
	for i := range grid {
		pessoa[i] = make([]int, len(grid[0]))
		fogo[i] = make([]int, len(grid[0]))
		for j := range grid[0] {
			pessoa[i][j] = -1
			fogo[i][j] = -1
		}
	}
	pessoa[0][0] = 0
	frente := []Path{{x: 0, y: 0, dia: 0}}
	for len(frente) > 0 {
		curr := frente[0]
		frente = frente[1:]
		for _, dir := range dirs {
			if curr.x+dir[0] >= 0 && curr.x+dir[0] < len(grid) && curr.y+dir[1] >= 0 && curr.y+dir[1] < len(grid[0]) && grid[curr.x+dir[0]][curr.y+dir[1]] == 0 && pessoa[curr.x+dir[0]][curr.y+dir[1]] == -1 {
				pessoa[curr.x+dir[0]][curr.y+dir[1]] = curr.dia + 1
				frente = append(frente, Path{x: curr.x + dir[0], y: curr.y + dir[1], dia: curr.dia + 1})
			}
		}
	}
	frenteFogo := make([]Path, 0)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				fogo[i][j] = 0
				frenteFogo = append(frenteFogo, Path{x: i, y: j, dia: 0})
			}
		}
	}
	for len(frenteFogo) > 0 {
		curr := frenteFogo[0]
		frenteFogo = frenteFogo[1:]
		for _, dir := range dirs {
			if curr.x+dir[0] >= 0 && curr.x+dir[0] < len(grid) && curr.y+dir[1] >= 0 && curr.y+dir[1] < len(grid[0]) && grid[curr.x+dir[0]][curr.y+dir[1]] == 0 && fogo[curr.x+dir[0]][curr.y+dir[1]] == -1 {
				fogo[curr.x+dir[0]][curr.y+dir[1]] = curr.dia + 1
				frenteFogo = append(frenteFogo, Path{x: curr.x + dir[0], y: curr.y + dir[1], dia: curr.dia + 1})
			}
		}
	}
	if pessoa[len(grid)-1][len(grid[0])-1] == -1 || fogo[len(grid)-1][len(grid[0])-1] < pessoa[len(grid)-1][len(grid[0])-1] {
		return -1
	}
	if fogo[len(grid)-1][len(grid[0])-1] == -1 {
		return 1000000000
	}
	diff := fogo[len(grid)-1][len(grid[0])-1] - pessoa[len(grid)-1][len(grid[0])-1]
	if pessoa[len(grid)-1][len(grid[0])-2] > -1 && pessoa[len(grid)-2][len(grid[0])-1] > -1 && (fogo[len(grid)-1][len(grid[0])-2]-pessoa[len(grid)-1][len(grid[0])-2] > diff || fogo[len(grid)-2][len(grid[0])-1]-pessoa[len(grid)-2][len(grid[0])-1] > diff) {
		return diff
	}
	return diff - 1
}

// Não modifique a função main
func main() {
	var nl, nc int
	fmt.Scan(&nl, &nc)
	grid := make([][]int, nl)
	for i := range grid {
		grid[i] = make([]int, nc)
	}
	for i := range grid {
		for j := range grid[0] {
			fmt.Scan(&grid[i][j])
		}
	}
	result := maximumMinutes(grid)
	fmt.Println(result)
}
