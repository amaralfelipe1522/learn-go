package main

import (
	"fmt"

	"github.com/amaralfelipe1522/learn-go/capitulo-19/02-packages/pkg2"
)

func main() {
	fmt.Println("***Exemplo de execução em pacotes***")
	Pkg1()
	result := Somar(1,2)
	fmt.Printf("Resultado da soma %d\n",result)

	result2 := pkg2.Multiplicar(result, 2)
	fmt.Printf("Resultado da multiplicação %d\n", result2)
}