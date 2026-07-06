package main

import "fmt"

var total int

func resolver(line, n int, col, d1, d2 []bool) {
	if line == n {
		total++
		return
	}
	for i := range n {
		if !col[i] && !d1[line+i] && !d2[line-i+n-1] {
			col[i] = true
			d1[line+i] = true
			d2[line-i+n-1] = true
			resolver(line+1, n, col,d1,d2)
			col[i] = false
			d1[line+i] = false
			d2[line-i+n-1] = false
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	col := make([]bool, n)
	d1 := make([]bool, 2*n)
	d2 := make([]bool, 2*n)
	total = 0
	resolver(0, n, col, d1, d2)
	fmt.Println(total)
}
