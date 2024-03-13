package main

import (
	"fmt"
)

// Buffer permite tirar sincronicidade do canal, não precisando que de imediato o valor seja usado em outro lugar
// Não é muito utilizado

func main() {

	canal := make(chan int, 1) // define o tipo do canal com buffer de qtd 1
	canal <- 1522

	printChannel(<-canal)

	// fmt.Println("Channel: ", <-canal) // causa deadlock pois ultrapassa a quantidade de buffer
}

func printChannel(n int) {
	fmt.Println("Channel: ", n)
}