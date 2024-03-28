package main

import (
	"fmt"
	"sync"
	"time"
)

// mesmo exemplo usando throttling (número maximo de goroutines)
// Abaixo ocorrem no maximo 5 go routines a cada 1 segundo
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	maximoPorVez := 5
	total := 100

	go func(x int) {
		for i:=0; i<x; i++ {
			c1 <- i
		}
		close(c1)
	}(total)
	
	go outra(maximoPorVez, c1, c2)

	for v := range c2 {
		fmt.Println("Processo na posição: ", v)
	}
}

func outra(qtd int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i:=0; i<qtd; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v) 
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
	
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n * 1000
}