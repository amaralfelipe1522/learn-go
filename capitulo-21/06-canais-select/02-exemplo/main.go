package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	x := 500

	go func(x int) {
		for i:=0; i<x; i++ {
			c <- i
			if (i == x-1) {
				quit <- i
			}
		}
	}(x/2)

	func() {
		for {
			select {
			case v := <-c:
				fmt.Printf("Valor %d veio do canal C\n",v)
			case v := <-quit:
				fmt.Printf("Valor %d veio do canal Quit\n",v)
				return
			}
		}
	}()
}