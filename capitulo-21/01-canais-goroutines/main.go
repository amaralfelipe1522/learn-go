package main

import (
	"fmt"
	"sync"
)

// Para cada entrada (var <-), uma saída (<-var)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	canal := make(chan int) // define o tipo do canal
	
	go func() {
		canal <- 1522
		wg.Done()
	}()
	// fmt.Println("Channel: ", <-canal)

	printChannel(<-canal)
	wg.Wait()
}

func printChannel(n int) {
	fmt.Println("Channel: ", n)
	wg.Done()
}