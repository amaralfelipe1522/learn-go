package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 1522
		close(c)
	}()

	v, ok := <- c
	fmt.Println(v, ok)

	v, ok = <- c
	fmt.Println(v, ok)

	// Se o valor inserido no canal fosse 0, sem o comma ok não daria pra saber se o valor foi de fato recebido do canal
}