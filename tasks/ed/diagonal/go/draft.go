package main
import "fmt"

func imprimir(s string, k int) {
    if len(s) == 0 {
        return
    }
    for range k {
        fmt.Print(" ")
    }
    fmt.Println(string(s[0]))
    imprimir(s[1:], k+1)
}

func main() {
    var palavra string
    fmt.Scan(&palavra)
    imprimir(palavra, 0)
}
