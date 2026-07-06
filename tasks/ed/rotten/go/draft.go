package main

import (
	"fmt"
	"slices"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l - 1, p.c},
		{p.l + 1, p.c},
		{p.l, p.c - 1},
		{p.l, p.c + 1},
	}
}

func inside(grid [][]int, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]int, pos Pos, num int) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == num
}

func cmpGrid(grid, backup [][]int) bool {
    if len(grid) != len(backup) {
        return false
    }
    for i := range grid {
        if !slices.Equal(grid[i], backup[i]) {
            return false
        }
    }
    return true
}

func updateGrid(grid, backup [][]int) {
    for l := range backup {
        for c := range backup[l] {
            if backup[l][c] == 2 {
                p := Pos{l: l, c: c}
                for _, neig := range p.getNeig() {
                    if match(backup, neig, 1) {
                        grid[neig.l][neig.c] = 2
                    }
                }
            }
        }
    }
}

func calcTime(grid [][]int) int {
    times := 0
    for {
        backup := make([][]int, len(grid))
        for i := range grid {
            backup[i] = make([]int, len(grid[i]))
            copy(backup[i], grid[i])
        }
        updateGrid(grid, backup)
        if cmpGrid(grid, backup) {
            break
        }
        times++
    }
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }
    return times
}

func main() {
    var m, n int
    fmt.Scan(&m, &n)
    grid := make([][]int, m)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for i := range grid {
        for j := range grid[i] {
            fmt.Scan(&grid[i][j])
        }
    }
    fmt.Println(calcTime(grid))
}
