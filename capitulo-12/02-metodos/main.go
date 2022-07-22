package main

import "fmt"

// Método é uma função anexada a um tipo

type Cachorro struct {
	nome  string
	idade int
}

func main() {

	shaman := Cachorro{
		"Shaman",
		13,
	}

	fmt.Printf("O %v tem %v anos e disse %v\n", shaman.nome, shaman.idade, shaman.latir())

}

func (c Cachorro) latir() string {
	return "Au Au"
}
