package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	x := 500

	go func(x int) {
		for i:=0; i<x; i++ {
			fmt.Println("Valor recebido: ", <-c) // sincronicamente adiciona numeroQualquer ao canal na outra func
		}
		quit <- 0
	}(x/2)

	func() {
		numeroQualquer := 1
		for { // for infinito até que o canal Quit receba um valor
			select {
			case c <- numeroQualquer:
				numeroQualquer++
			case <-quit:
				return
			}
		}
	}()
}