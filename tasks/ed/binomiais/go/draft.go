package main
import "fmt"

func binomio(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    return binomio(n-1, k-1) + binomio(n-1, k)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    fmt.Println(binomio(n, k))
}
