package main
import "fmt"

func pode(linha []byte, l, pos, valor int) bool {
    for i := 1; i <= l; i++ {
        if pos+i < len(linha) && linha[pos+i] == byte(valor) + '0' || pos-i >= 0 && linha[pos-i] == byte(valor) + '0' {
            return false
        }
    }
    return true
}

func aux(linha []byte, l, pos int) bool {
    if pos == len(linha) {
        return true
    }
    if linha[pos] != '.' {
        return aux(linha, l, pos+1)
    }
    for i := range l+1 {
        if pode(linha, l, pos, i) {
            linha[pos] = byte(i + '0')
            if aux(linha, l, pos+1) {
                return true
            }
            linha[pos] = '.'
        }
    }
    return false
}

func processa(linha string, l int) string {
    novalinha := []byte(linha)
    aux(novalinha, l, 0)
    return string(novalinha)
}

func main() {
    var linha string
    var l int
    fmt.Scan(&linha, &l)
    fmt.Println(processa(linha, l))
}