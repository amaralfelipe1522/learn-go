package main

import "fmt"

func main() {

	fmt.Println(somaOuSubtrai(11)(11))
	fmt.Println(somaOuSubtrai(8)(8))

}

func somaOuSubtrai(value int) func(v int) int {
	if value > 10 {
		return func(v int) int {
			return v - 2
		}
	} else {
		return func(v int) int {
			return v + 2
		}
	}
}
