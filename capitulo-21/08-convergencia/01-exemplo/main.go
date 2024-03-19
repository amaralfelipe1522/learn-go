package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebido(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido no Converge", v)
	}
}

func envia(p, i chan int) {
	total := 100

	for j:=0; j<total; j++ {
		if j % 2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	defer close(p)
	defer close(i)
}

func recebido(p, i, c chan int) {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)
}