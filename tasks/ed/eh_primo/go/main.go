package main

import (
	"fmt"
	"math"
)

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x == 2 {
		return true
	}
	if x % div == 0 || x == 1 {
		return false
	}
	if div < int(math.Sqrt(float64(x))) {
		eh_primo(x, div+1)
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
