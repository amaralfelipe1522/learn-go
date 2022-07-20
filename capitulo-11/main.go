package main

import "fmt"

type Pessoa struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {
	// Exercício 1
	pessoa1 := Pessoa{
		"Carol",
		"Carles",
		[]string{"Flocos", "Chocolate"},
	}
	fmt.Printf("%v %v gosta dos sorvetes:\n", pessoa1.nome, pessoa1.sobrenome)
	for _, sorvete := range pessoa1.sorvetesFavoritos {
		fmt.Printf("\t%v\n", sorvete)
	}

	pessoa2 := Pessoa{
		"Celta",
		"Luiz",
		[]string{"Creme", "Uva"},
	}
	fmt.Printf("%v %v gosta dos sorvetes:\n", pessoa2.nome, pessoa2.sobrenome)
	for _, sorvete := range pessoa1.sorvetesFavoritos {
		fmt.Printf("\t%v\n", sorvete)
	}

	// Exercício 2
	mapPessoas := map[string]Pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}
	for _, pessoa := range mapPessoas {
		fmt.Printf("%v %v gosta dos sorvetes:\n", pessoa.nome, pessoa.sobrenome)
		for _, sorvete := range pessoa.sorvetesFavoritos {
			fmt.Printf("\t%v\n", sorvete)
		}
	}
}
