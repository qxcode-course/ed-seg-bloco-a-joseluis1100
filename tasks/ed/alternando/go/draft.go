// teste á
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Soldado struct {
	vivo, positivo bool
}

func iniciarSoldados(N, F int) []Soldado {
	soldados := make([]Soldado, N)
	for i := range soldados {
		if F == 1 && i%2 == 0 || F == -1 && i%2 != 0 {
			soldados[i].positivo = true
		}
		soldados[i].vivo = true
	}
	return soldados
}

func imprimirSoldados(E int, soldados []Soldado) {
	vetor := make([]string, 0)
	for i := range soldados {
		if i == E && soldados[i].positivo {
			vetor = append(vetor, strconv.Itoa(i+1)+">")
		} else if i == E && !soldados[i].positivo {
			vetor = append(vetor, "<-"+strconv.Itoa(i+1))
		} else if soldados[i].vivo && soldados[i].positivo {
			vetor = append(vetor, strconv.Itoa(i+1))
		} else if soldados[i].vivo && !soldados[i].positivo {
			vetor = append(vetor, "-"+strconv.Itoa(i+1))
		}
	}
	fmt.Printf("[ %s ]\n", strings.Join(vetor, " "))
}

func decidirAlvo(operacao int, pos int, soldados []Soldado) int {
	for i, j := ((pos+operacao)%len(soldados)+len(soldados))%len(soldados), 0; j < len(soldados); i, j = ((i+operacao)%len(soldados)+len(soldados))%len(soldados), j+1 {
		if soldados[i].vivo {
			return i
		}
	}
	return -1
}

func jogo(E int, soldados []Soldado) {
	var positividade int
	if soldados[E].positivo {
		positividade = 1
	} else {
		positividade = -1
	}
	alvo := decidirAlvo(positividade, E, soldados)
	if alvo == -1 {
		return
	}
	imprimirSoldados(E, soldados)
	soldados[alvo].vivo = false
	alvo = decidirAlvo(positividade, E, soldados)
	if alvo == -1 {
		return
	}
	jogo(decidirAlvo(positividade, E, soldados), soldados)
}

func main() {
	var N, E, F int
	fmt.Scan(&N, &E, &F)
	soldados := iniciarSoldados(N, F)
	jogo(E-1, soldados)
}
