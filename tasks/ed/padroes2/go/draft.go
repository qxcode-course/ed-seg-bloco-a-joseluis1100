package main
import "fmt"

func calcular(n int) int {
    resultado := n*n+2*n
    return resultado
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(calcular(n))
}
