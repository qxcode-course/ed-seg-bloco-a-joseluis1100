package main

import "fmt"

func fuga(H, P, F, D int) string {
	for {
		F = (F + D + 16) % 16
		if F == H {
			return "S"
		}
		if F == P {
			return "N"
		}
	}
}

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)
	fmt.Println(fuga(H, P, F, D))
}
