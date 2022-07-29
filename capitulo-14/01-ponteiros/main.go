package main

import "fmt"

func main() {
	x := 0
	fmt.Printf("O valor de X é %v e está no endereço %v\n", x, &x)
	y := &x
	fmt.Printf("O valor de Y é %v e está no endereço %v\n", *y, y)
	*y++
	fmt.Printf("O valor de X é %v e está no endereço %v\n", x, &x)
}
