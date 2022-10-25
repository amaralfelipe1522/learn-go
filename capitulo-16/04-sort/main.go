package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPotencia []carro

func (p ordenarPotencia) Len() int {
	return len(p)
}

func (p ordenarPotencia) Less(i, j int) bool {
	return p[i].potencia > p[j].potencia
}

func (p ordenarPotencia) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	carros := []carro{
		{
			"Celta",
			78,
			10,
		}, {
			"Vectra A",
			151,
			9,
		},
	}

	fmt.Println(carros)

	sort.Sort(ordenarPotencia(carros))

	fmt.Println(carros)

}
