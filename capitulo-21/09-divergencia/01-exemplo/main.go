package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	x := 100
	go func(x int) {
		for i:=0; i<x; i++ {
			c1 <- i
		}
		close(c1)
	}(x)
	
	go outra(c1, c2)

	for v := range c2 {
		fmt.Println("Processo na posição: ", v)
	}
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup
	// Diverge o trabalho em N GO Routines (usar throttling para não sobrecarregar o processador no próximo exemplo)
	for v := range canal1 {
		wg.Add(1)
		go func (v int)  {
			canal2 <- trabalho(v) // Converge todo o trabalho em um unico canal
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n * 1000
}