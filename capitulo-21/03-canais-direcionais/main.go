package main

import "fmt"

func main() {
	canal := make(chan int) // criado de forma bidirecional, porém, poderia ser por exemplo make(<-chan int)

	go apenasRecebeValor(canal)

	apenasExibeValor(canal)
}

func apenasRecebeValor(r chan<- int) {
	r <- 1522
}

func apenasExibeValor(e <-chan int) {
	fmt.Println("Valor para exibir: ", <-e)
}