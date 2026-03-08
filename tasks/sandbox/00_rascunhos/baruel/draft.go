package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adicionarFigurinhas(quantidade int, figurinhas map[int]int, repetidas *[]string) {
    for range quantidade {
        var figurinha int
        fmt.Scan(&figurinha)
        figurinhas[figurinha]++
        if figurinhas[figurinha] > 1 {
            *repetidas = append(*repetidas, strconv.Itoa(figurinha))
        }
    }
}

func encontrarFaltando(total int, figurinhas map[int]int, faltando *[]string) {
    for i := 1; i <= total; i++ {
        if figurinhas[i] == 0 {
            *faltando = append(*faltando, strconv.Itoa(i))
        }
    }
}

func resultados(vetor []string) {
    if len(vetor) == 0 {
        fmt.Println("N")
        return
    }
    fmt.Println(strings.Join(vetor, " "))
}

func main() {
    var total, quantidade int
    fmt.Scan(&total, &quantidade)
    figurinhas := make(map[int]int)
    repetidas := make([]string, 0)
    faltando := make([]string, 0)
    adicionarFigurinhas(quantidade, figurinhas, &repetidas)
    encontrarFaltando(total, figurinhas, &faltando)
    resultados(repetidas)
    resultados(faltando)
}
