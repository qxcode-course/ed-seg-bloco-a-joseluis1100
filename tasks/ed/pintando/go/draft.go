package main

import (
	"fmt"
	"math"
)

func heron(a float64, b float64, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}
func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%.2f\n", heron(a, b, c))
}
