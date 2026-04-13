package main

import (
	"fmt"
	"math"
)

func ativos(n, k, partes int) int {
    if n <= k {
        return 0
    }
    if n % 2 == 0 {
        partes += 2
        p1 = ativos(n/2, k, partes)
    } else {
        partes += 4
        p1 = ativos(int(math.Ceil(float64(n)/2)), k, partes)
        p2 = ativos(int(math.Floor(float64(n)/2)), k, partes)
    }
    return partes
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    fmt.Println(ativos(n, k, 0))
}
