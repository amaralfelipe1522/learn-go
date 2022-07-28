package main

import "fmt"

// Closure é cercar ou capturar um scope para que possamos utiliza-lo em outro contexto
func main() {
	a := incrementar()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := incrementar()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementar() func() int {
	x := 0 // O scope do X será diferente a cada cópia da função Closure
	return func() int {
		x++
		return x
	}
}
