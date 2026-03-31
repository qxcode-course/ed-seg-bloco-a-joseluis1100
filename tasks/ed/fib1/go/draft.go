package main
import "fmt"

func coelhos(n, k int) int {
    if n == 1 || n == 2 {
        return 1
    }
    pares := coelhos(n-1, k) + coelhos(n-2, k)
    return pares
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    fmt.Println(coelhos(n, k))
}
