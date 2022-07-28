package main

import "fmt"

// Exemplo de recursividade com fatorial 4! (4 * 3 * 2 * 1 = 24):

func main() {
	valor := 52
	result, _ := fatorial(valor, valor)
	fmt.Printf("O fatorial de %v é %v\n", valor, result)

}

func fatorial(resultado, valorInicial int) (int, int) {
	if valorInicial == 1 {
		return resultado, valorInicial
	}
	mult := resultado * (valorInicial - 1)
	valorInicial--
	return fatorial(mult, valorInicial)
}
