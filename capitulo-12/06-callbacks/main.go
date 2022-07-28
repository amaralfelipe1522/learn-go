package main

import "fmt"

// Callback se trata de passar uma função como argumento de outra função
func main() {
	fmt.Println(verificaNumerosPares(somaApenasNumerosPares, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
}

func verificaNumerosPares(s func(pares ...int) int, sliceInt ...int) int {
	var numerosPares []int
	for _, par := range sliceInt {
		if par%2 == 0 {
			numerosPares = append(numerosPares, par)
		}
	}

	total := s(numerosPares...)

	return total
}

func somaApenasNumerosPares(slicePares ...int) int {
	var total int
	for _, value := range slicePares {
		total += value
	}
	return total
}
