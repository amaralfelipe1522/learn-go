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

	// Exercício 3
	type Veiculo struct {
		portas int
		cor    string
	}

	type Caminhonete struct {
		Veiculo
		quatroRodas bool
	}

	type Sedan struct {
		Veiculo
		modeloLuxo bool
	}

	pajero := Caminhonete{
		Veiculo{4, "azul"},
		true,
	}

	fmt.Printf("Modelo da caminhonete possui %v portas e é da cor %v\n", pajero.portas, pajero.cor)

	vectra := Sedan{
		Veiculo{4, "preto"},
		true,
	}

	fmt.Printf("Modelo do sedan possui %v portas e é da cor %v\n", vectra.portas, vectra.cor)

	// Exercício 4
	compras := struct {
		local   map[string]int
		compras []string
	}{
		map[string]int{
			"Valor": 200,
		},
		[]string{"Encosto de cama", "ventilador de teto"},
	}

	fmt.Println(compras.compras[0])
	fmt.Println(compras.local["Valor"])
}
