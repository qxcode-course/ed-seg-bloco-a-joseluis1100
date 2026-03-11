package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var primos = []string{}

func eh_primo(x int, div int) bool {
	if x == 2 {
		return true
	}
	if x % div == 0 || x == 1 {
		return false
	}
	if div < int(math.Sqrt(float64(x))) {
		return eh_primo(x, div+1)
	}
	return true
}

func procurarPrimo(n, valor int) {
    if n == 0 {
        return
    }
    for {
        if eh_primo(valor, 2) {
            primos = append(primos, strconv.Itoa(valor))
            break
        } else {
            valor++
        }
    }
    procurarPrimo(n-1, valor+1)
}

func main() {
    var n int
    fmt.Scan(&n)
    procurarPrimo(n, 2)
    fmt.Printf("[%s]\n", strings.Join(primos, ", "))
}
