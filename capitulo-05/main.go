package main

import "fmt"

func main() {
	// Exercício 1
	num := 31337
	fmt.Printf("%d\n%b\n%#x\n", num, num, num)

	num2 := num << 1
	fmt.Printf("%d\n%b\n%#x\n", num2, num2, num2)

	rawString := `Exercício\t2\n`
	fmt.Println(rawString)

	// Exercício 4
	const (
		ano1 = iota + 2023
		ano2
		ano3
		ano4
	)

	fmt.Printf("%v\n%v\n%v\n%v\n", ano1, ano2, ano3, ano4)
}
