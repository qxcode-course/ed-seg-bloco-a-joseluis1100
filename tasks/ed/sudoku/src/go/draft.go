package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func pode(matriz [][]rune, quad []rune, lin, col, valor int) bool {
    if slices.Contains(quad, rune(valor) + '0') {
            return false
        }
    for i := range matriz {
        if matriz[i][col] == rune(valor) + '0' {
            return false
        }
    }
    for i := range matriz {
        if matriz[lin][i] == rune(valor) + '0' {
            return false
        }
    }
    return true
}

func quadrante(matriz [][]rune, lin, col int) []rune {
    dim := len(matriz)

    if dim == 4 {
        l := int(lin / 2) * 2
        c := int(col / 2) * 2
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1],
            matriz[l+1][c], matriz[l+1][c+1],
        }
    }

    if dim == 9 {
        l := int(lin / 3) * 3
        c := int(col / 3) * 3
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1], matriz[l+0][c+2],
            matriz[l+1][c], matriz[l+1][c+1], matriz[l+1][c+2],
            matriz[l+2][c], matriz[l+2][c+1], matriz[l+2][c+2],
        }
    }
    return nil
}

func resolver(matriz [][]rune, index int) bool {
    nl := len(matriz)
    l := index / nl
    c := index % nl
    if index == nl * nl {
        return true
    }
    // se não for ponto, continue
    if matriz[l][c] != '.' {
        return resolver(matriz, index+1)
    }
    // para todos os números de [1 a N]
    for i := 1; i <= nl; i++ {
    //     se o número não estiver na linha, coluna e quadrante
        if pode(matriz, quadrante(matriz, l, c), l, c, i) {
    //         coloque o número na matriz
            matriz[l][c] = rune(i) + '0' 
    //         se resolver(matriz, index + 1):
            if resolver(matriz, index+1) {
    //             return True
                return true
            }
    //         matriz[l][c] = '.' // desfaz a tentativa
            matriz[l][c] = '.'
        }
}
    // return False
    return false
}

func main() {
    var n int
    fmt.Scan(&n)
    matriz := make([][]rune, 0)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		matriz = append(matriz, []rune(scanner.Text()))
	}
    resolver(matriz, 0)
    for i := range matriz {
        for j := range matriz[0] {
            fmt.Printf("%c", matriz[i][j])
        }
        fmt.Println()
    }
}