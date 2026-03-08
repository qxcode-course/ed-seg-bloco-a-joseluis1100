package main

import (
	"fmt"
	"math"
)

func ler(x *float64) {
	fmt.Scan(x)
}
func heron(a float64, b float64, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}
func main() {
	var a, b, c float64
	ler(&a)
	ler(&b)
	ler(&c)
	fmt.Printf("%.2f\n", heron(a, b, c))
}
