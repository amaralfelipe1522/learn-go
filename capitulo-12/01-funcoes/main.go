package main

import "fmt"

func main() {
	defer fmt.Println(fatorial(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52))

	// Ou colocar os valores em uma slice e desenrola-las
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	fmt.Println(fatorial(sliceInt...))

	// Funções com parametros variaticos podem receber 0 como argumento na chamada
	fmt.Println(fatorial())
}

func fatorial(x ...int) uint64 {
	soma := 0

	for _, valor := range x {
		if soma == 0 {
			soma = valor
		} else {
			soma *= valor
		}
	}

	return uint64(soma)
}
