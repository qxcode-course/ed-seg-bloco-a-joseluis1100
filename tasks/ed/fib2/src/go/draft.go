package main
import "fmt"

func coelhos(n int) int {
    if n <= 2 {
        return 1
    }
    if n <= 4 {
        return 2
    }
    pares := coelhos(n-1) + coelhos(n-2) - coelhos(n-4)
    return pares
}

func main() {
    var mes int
    fmt.Scan(&mes)
    fmt.Println(coelhos(mes))
}
