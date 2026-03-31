package main
import "fmt"

func calcular(n int) {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return
    }
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, n-1, n-1)
    calcular(n-1)
}

func reverso(i, n int) {
    if i == n + 1 {
        return
    }
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", i, i-1, i-1, i * i)
    reverso(i+1, n)
}

func main() {
    var n int
    fmt.Scan(&n)
    calcular(n)
    reverso(2, n)
}
