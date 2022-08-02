package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	// Exercício
	pessoa := pessoa{
		"Felipe",
		30,
	}

	fmt.Println(pessoa.aniversario())
	fmt.Println(pessoa.idade)
}

func (p *pessoa) aniversario() int {
	(*p).idade++
	return p.idade
}
