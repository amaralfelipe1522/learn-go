package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c1 := trabalho("Sprint 1")
	c2 := trabalho("Sprint 2")
	c3 := trabalho("Sprint 3")


	wg.Add(1)
	convergido := converge(c1, c2, c3)
	
	for i:=0; i<50; i++{
		fmt.Println(<-convergido)
	}

	wg.Done()
	wg.Wait()
}

func trabalho(s string) chan string {
	canal := make(chan string)
	
	go func(s string, c chan string) {
		for i:=0; ; i++ {
			canal <- fmt.Sprintf("Função %s realizada -> %d\n", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3))) // duração entre 0 e 1000 milisegundos
		}
	}(s, canal)

	return canal
}

func converge(x, y, z chan string) chan string {
	converge := make(chan string)

	go func(x chan string) {
		for {
			converge <- <- x
		}
	}(x)

	go func(y chan string) {
		for {
			converge <- <- y
		}
	}(y)

	go func(z chan string) {
		for {
			converge <- <- z
		}
	}(z)

	return converge
}