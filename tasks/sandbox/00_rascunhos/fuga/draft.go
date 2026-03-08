package main

import "fmt"

func ler(x *int) {
	fmt.Scan(x)
}
func main() {
	var H, P, F, D int
    ler(&H)
    ler(&P)
    ler(&F)
    ler(&D)
	for {
		F = (F + D + 16) % 16
		if F == H {
			fmt.Println("S")
			return
		}
		if F == P {
			fmt.Println("N")
			return
		}
	}
}
