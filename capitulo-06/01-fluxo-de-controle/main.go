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
}
