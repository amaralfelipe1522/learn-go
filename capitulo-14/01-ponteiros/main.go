package main

import "fmt"

type numbers struct {
	first  int
	second int
}

func main() {
	x := 0
	fmt.Printf("O valor de X é %v e está no endereço %v\n", x, &x)
	y := &x
	fmt.Printf("O valor de Y é %v e está no endereço %v\n", *y, y)
	*y++
	fmt.Printf("O valor de X é %v e está no endereço %v\n", x, &x)

	numbers := numbers{
		2,
		4,
	}

	fmt.Println(numbers.somarPoint())
	fmt.Println(numbers.somarPoint())
	fmt.Println(numbers.somarPoint())
	fmt.Println(numbers.somar())
	fmt.Println(numbers.somar())
	fmt.Println(numbers.somarPoint())
}

func (n numbers) somar() (int, int) {
	n.first++
	n.second++
	return n.first, n.second
}

func (n *numbers) somarPoint() (int, int) {
	n.first++
	n.second++
	return n.first, n.second
}
