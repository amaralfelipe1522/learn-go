package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go send(par, impar, quit)
	receive(par, impar, quit)
}

func send(par, impar chan int, quit chan bool) {
	total := 200
	for i:=0; i<total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	defer close(par)
	defer close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v, ok := <- par:
			if (!ok) {
				return
			} else {
				fmt.Printf("Valor %d recebido no canal Par. Comma OK: %t\n", v, ok)
			}
		case v, ok := <- impar:
			if (!ok) {
				return
			} else {
				fmt.Printf("Valor %d recebido no canal Impar. Comma OK: %t\n", v, ok)
			}
		case <-quit:
			fmt.Printf("Processo encerrado\n")
			return
		}
	}
}