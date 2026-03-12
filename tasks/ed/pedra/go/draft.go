package main

import (
	"fmt"
	"math"
	"strconv"
)

type Pessoa struct {
	A, B         int
	classificado bool
}

func ler(x *int) {
	fmt.Scan(x)
}

func adicionarPessoa() Pessoa {
	var pessoa Pessoa
	ler(&pessoa.A)
	ler(&pessoa.B)
	if pessoa.A < 10 || pessoa.B < 10 {
		pessoa.classificado = false
	} else {
		pessoa.classificado = true
	}
	return pessoa
}

func atualizarMelhor(pessoa1 Pessoa, pessoa2 Pessoa) bool {
	if math.Abs(float64(pessoa1.A-pessoa1.B)) < math.Abs(float64(pessoa2.A-pessoa2.B)) {
		return true
	}
	return false
}

func definirMelhor(repeticoes int, pessoas []Pessoa) string {
	melhor := -1
	for i := range repeticoes {
		pessoas[i] = adicionarPessoa()
		if melhor == -1 && pessoas[i].classificado {
			melhor = i
		}
		if pessoas[i].classificado && atualizarMelhor(pessoas[i], pessoas[melhor]) {
			melhor = i
		}
	}
	if melhor == -1 {
		return "sem ganhador"
	} else {
		return strconv.Itoa(melhor)
	}
}

func main() {
	var repeticoes int
	ler(&repeticoes)
	pessoas := make([]Pessoa, repeticoes)
	fmt.Println(definirMelhor(repeticoes, pessoas))
}
