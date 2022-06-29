package main

import (
	"fmt"
	"log"
)

var packageScope = true

func main() {
	bts, err := fmt.Println("Hello World")
	fmt.Println(bts)
	if err != nil {
		log.Fatalf("Error")
		fmt.Println(err)
	}

	x := "string"
	y := 2

	fmt.Printf("x: %v - %T\n", x, x)
	fmt.Printf("y: %v - %T\n", y, y)
	y = 3
	fmt.Printf("y: %v - %T\n", y, y)
	y, z := 4, true
	fmt.Printf("y: %v - %T\n", y, y)
	fmt.Printf("z: %v - %T\n", z, z)

	// Criando próprio tipo
	type meuInt int // int como tipo subjacente

	var num1 meuInt
	fmt.Printf("num1: %v - tipo: %T\n", num1, num1)

	// Conversão de tipos
	var num2 = int(num1)
	fmt.Printf("num2: %v - tipo: %T\n", num2, num2)

}
