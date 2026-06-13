package main
import "fmt"

func pode(linha string, l, pos, valor int) bool {
    for i := range l {
        if linha[pos+i] == byte(valor) || linha[pos-i] == byte(valor) {
            return false
        }
    }
    return true
}

func processa(linha string, l int) string {
    novalinha := []byte(linha)
    for i := range linha {
        if linha[i] == '.' {
            for j := 0; j <= l; j++ {
                if pode(linha, l, i, j){
                    novalinha[i] = byte(j + '0')
                    break
                }
            }
        }
    }
    return string(novalinha)
}

func main() {
    var linha string
    var l int
    fmt.Scan(linha, l)
    println(processa(linha, l))
}