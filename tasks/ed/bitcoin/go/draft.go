package main

import (
	"fmt"
	"math"
)

func ativos(n, k int) int {
    if n <= k {
        return 1
    }
    return ativos(int(math.Ceil(float64(n)/2)), k) + ativos(int(math.Floor(float64(n)/2)), k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    fmt.Println(ativos(n, k))
}
