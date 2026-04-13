package main
import "fmt"

func contaPares(n int) int {
    var aux, pares int
    mapa := make(map[int]int)
    for range n {
        fmt.Scan(&aux)
        if mapa[-aux] != 0 {
            mapa[-aux]--
            pares++
        } else {
            mapa[aux]++
        }
    }
    return pares
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(contaPares(n))
}
