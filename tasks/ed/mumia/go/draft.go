package main

import "fmt"

func estado(idade int) string {
	switch {
	case idade < 12:
		return "crianca"
	case idade < 18:
		return "jovem"
	case idade < 65:
		return "adulto"
	case idade < 1000:
		return "idoso"
	default:
		return "mumia"
	}
}

func main() {
	var nome string
	var idade int
	fmt.Scan(&nome, &idade)
	fmt.Printf("%s eh %s\n", nome, estado(idade))
}
