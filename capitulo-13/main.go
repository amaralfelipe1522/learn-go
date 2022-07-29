package main

import (
	"fmt"
	"math/rand"
	"time"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	// Exercício 1
	fmt.Println(retornaInt())
	fmt.Println(retornaIntString())

	// Exercício 2
	fmt.Println(retornaSomaVariatico(1, 2, 3, 4, 5))
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(retornaSomaSlice(slice))

	// Exercício 3
	func() {
		defer fmt.Println("-> 1")
		fmt.Println("-> 2")
		fmt.Println("-> 3")
	}()

	// Exercício 4
	pessoa := pessoa{
		"Celta",
		"Tolansky",
		34,
	}
	pessoa.exibePessoa()

	// Exercício 7
	funcAnon := func() {
		defer fmt.Println("-> 1")
		fmt.Println("-> 2")
		fmt.Println("-> 3")
	}
	funcAnon()

	// Exercício 8
	retorno := retornaFunc()
	retorno()

	// Exercício 9
	fmt.Println(multiplicaInt(retornaInt))

	// Exercício 10
	a := closureExample()
	fmt.Println("Escopo de A: ", a())
	fmt.Println("Escopo de A: ", a())
	fmt.Println("Escopo de A: ", a())

	b := closureExample()
	fmt.Println("Escopo de B: ", b())
	fmt.Println("Escopo de B: ", b())
	fmt.Println("Escopo de B: ", b())
}

func retornaInt() int {
	x := rand.NewSource(time.Now().UnixNano())
	y := rand.New(x)
	number := y.Intn(100)
	return number
}

func retornaIntString() (int, string) {
	x := rand.NewSource(time.Now().UnixNano())
	y := rand.New(x)
	number := y.Intn(100)
	stringNumber := fmt.Sprint(y.Intn(100))
	return number, stringNumber
}

func retornaSomaVariatico(x ...int) int {
	var result int
	for _, x := range x {
		result += x
	}

	return result
}

func retornaSomaSlice(x []int) int {
	var result int
	for _, x := range x {
		result += x
	}

	return result
}

func (p pessoa) exibePessoa() {
	fmt.Printf("O %v %v tem %v anos\n", p.nome, p.sobrenome, p.idade)
}

func retornaFunc() func() {
	return func() {
		fmt.Println("Retornando a função dentro de outra função")
	}
}

func multiplicaInt(f func() int) int {
	soma := f()
	return soma * soma
}

func closureExample() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
