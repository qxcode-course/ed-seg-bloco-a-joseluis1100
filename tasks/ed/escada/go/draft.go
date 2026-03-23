package main
import "fmt"

func escadas(n int) int {
    if n == 1 || n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    return escadas(n-1) + escadas(n-3)
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(escadas(n))
}
