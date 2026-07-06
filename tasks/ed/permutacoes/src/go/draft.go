package main

import (
	"fmt"
	"slices"
)

func sortString(line string) string {
    chars := []byte(line)
    slices.Sort(chars)
    return string(chars)
}

func aux(resultados *[]string, resto, atual string) {
    if resto == "" {
        *resultados = append(*resultados, atual)
        return
    }
    for i := range resto {
        aux(resultados, resto[:i]+resto[i+1:], atual+string(resto[i]))
    }
}

func perm(line string) []string {
    line = sortString(line)
    var resultados []string
    aux(&resultados, line, "")
    return resultados
}

func main() {
    var line string
    fmt.Scan(&line)
    for _, perm := range perm(line) {
        fmt.Println(perm)
    }
}