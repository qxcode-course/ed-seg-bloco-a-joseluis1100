package main

import (
	"fmt"
	"math"
)

func maximumDetonation(bombs [][]int) int {
    neig := make(map[int][]int)
    for i := range bombs {
        for j := range bombs {
            if i != j && math.Pow(float64(bombs[j][0]-bombs[i][0]), 2)+math.Pow(float64(bombs[j][1]-bombs[i][1]), 2) <= math.Pow(float64(bombs[i][2]), 2) {
                neig[i] = append(neig[i], j)
            }
        }
    }
    var max int
    for i := range bombs {
        visited := make(map[int]bool)
        queue := []int{i}
        visited[i] = true
        for len(queue) > 0 {
            tmp := queue[0]
            queue = queue[1:]
            if _, ok := neig[tmp]; ok {
                for j := range neig[tmp] {
                    if !visited[neig[tmp][j]] {
                        visited[neig[tmp][j]] = true
                        queue = append(queue, neig[tmp][j])
                    }
                }
            }
        }
        if len(visited) == len(bombs) {
            return len(visited)
        }
        if len(visited) > max {
            max = len(visited)
        }
    }
    return max
}

func main() {
    var l, c int
    fmt.Scan(&l, &c)
    bombs := make([][]int, l)
    for i := range l {
        bombs[i] = make([]int, c)
    }
    for i := range bombs {
        for j := range bombs[0] {
            fmt.Scan(&bombs[i][j])
        }
    }
    fmt.Println(maximumDetonation(bombs))
}