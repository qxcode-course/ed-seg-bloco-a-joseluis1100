package main
import "fmt"

func calcular(n int) int {
    resultado := ((2+n)*8)-4
    return resultado
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(calcular(n))
}
