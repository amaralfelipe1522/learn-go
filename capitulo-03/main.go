package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

var y2 int

func main() {
	// Exercício 1
	// fmt.Println("Exercício 1")
	// x := 42
	// y := "James Bond"
	// z := true

	// fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// Exercício 2
	fmt.Println("Exercício 2")

	s := fmt.Sprintf("s: %v%v%v", x, y, z)
	fmt.Println(s)

	// Exercício 4
	fmt.Println("Exercício 4")
	type meuInt int
	var x meuInt
	fmt.Printf("x tem o valor %v e é do tipo %T\n", x, x)
	x = 42
	fmt.Printf("Agora x tem o valor %v e é do tipo %T\n", x, x)

	// Exercício 5
	fmt.Println("Exercício 5")
	y2 = int(x)
	fmt.Printf("y2 tem o valor %v e é do tipo %T\n", y2, y2)

}
