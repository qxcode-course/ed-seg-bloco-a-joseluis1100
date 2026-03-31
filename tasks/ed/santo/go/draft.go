package main
import "fmt"

func milagre(n, c, valor float64) {
    if n == 1 {
        fmt.Printf("%.2f\n", valor/2)
        return
    }
    valor = valor/2 + c
    milagre(n-1, c, valor)
}

func main() {
    var n, c float64
    fmt.Scan(&n, &c)
    milagre(n, c, c)
}
