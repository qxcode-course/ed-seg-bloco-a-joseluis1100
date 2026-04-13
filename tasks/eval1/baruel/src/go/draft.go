package main

import (
	"fmt"
	"strconv"
	"strings"
)

func verifica(total, possui int) ([]int, []int) {
    var aux int
    mapa := make(map[int]int)
    repetidas := make([]int, 0)
    faltando := make([]int, 0)
    for range possui {
        fmt.Scan(&aux)
        mapa[aux]++
    }
    for i := range total + 1 {
        if mapa[i] > 1 {
            for range mapa[i] -1 {
                repetidas = append(repetidas, i)
            }
        } else if mapa[i] == 0 && i != 0 {
            faltando = append(faltando, i)
        }
    }
    return repetidas, faltando
}

func toString(vetor []int) string {
    if len(vetor) == 0 {
        return "N"
    }
    vetString := make([]string, 0)
    for _, valor := range vetor {
        vetString = append(vetString, strconv.Itoa(valor)) 
    }
    return strings.Join(vetString, " ")
}

func main() {
    var total, possui int
    fmt.Scan(&total, &possui)
    repetidas, faltando := verifica(total, possui)
    fmt.Println(toString(repetidas))
    fmt.Println(toString(faltando))
}
