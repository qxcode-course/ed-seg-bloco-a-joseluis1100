package main

import (
	"fmt"
	"slices"
)

type Gomo struct {
	x, y int
}

var movimentos = map[string]Gomo{
    "U": {0, -1},
    "D": {0, 1},
    "L": {-1, 0},
    "R": {1, 0},
}

func escanearGomos(cobra []Gomo) {
	for i := range cobra {
		fmt.Scan(&cobra[i].x, &cobra[i].y)
	}
}

func avancar(direcao string, cobra []Gomo) []Gomo {
    mudaX, mudaY := movimentos[direcao].x, movimentos[direcao].y
	cobra = slices.Insert(cobra, 0, Gomo{x: cobra[0].x + mudaX, y: cobra[0].y + mudaY})
	return slices.Delete(cobra, len(cobra)-1, len(cobra))
}

func imprimirSlice(cobra []Gomo) {
	for _, gomo := range cobra {
		fmt.Printf("%d %d\n", gomo.x, gomo.y)
	}
}

func main() {
	var quantidade int
	var direcao string
	fmt.Scan(&quantidade, &direcao)
	cobra := make([]Gomo, quantidade)
	escanearGomos(cobra)
	imprimirSlice(avancar(direcao, cobra))
}