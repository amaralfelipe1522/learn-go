package main

import "fmt"

// 3 tipos de fluxo de controle: Sequencial, Repetição e Condicional
func main() {
	// Função hierarquica
	for x := 0; x <= 2; x++ {
		fmt.Printf("Minutos: %02d:00\n", x)
		for y := 1; y < 60; y++ {
			fmt.Printf("- Segundos: %02d:%02d\n", x, y)
			if y == 58 {
				break
			}
		}
	}

	// Exemplo de 'continue'
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Número par encontrado: %d\n", i)
	}

	// Utilizando ASCII
	for num := 33; num <= 122; num++ {
		fmt.Printf("Decimal: %d\tHexadecimal: %#x\tUnicode: %#U\n", num, num, num)
	}

	num := 8

	switch {
	case num < 10:
		fmt.Println("Menor que 10")
		fallthrough
	case num < 50:
		fmt.Println("Menor que 50")
	default:
		fmt.Println("Maior que 50")
	}

}
