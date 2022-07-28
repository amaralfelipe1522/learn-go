package main

import "fmt"

func main() {
	x := 4
	y := 6

	func(n1 int, n2 int) {
		n3 := n1 + n2
		fmt.Printf("A soma de %v e %v é %v\n", n1, n2, n3)
	}(x, y)

	// função como expressão

	z := func(n1 int, n2 int) {
		n3 := n1 + n2
		fmt.Printf("A soma de %v e %v é %v\n", n1, n2, n3)
	}

	z(x, y)
}
