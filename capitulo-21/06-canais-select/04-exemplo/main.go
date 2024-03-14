package main

import "fmt"

// Este tem um problema, onde se retorna o Case de Valor 0 mais de uma vez
// Solução no próximo exemplo

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
	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <- par:
			fmt.Printf("Valor %d recebido no canal Par\n", v)
		case v := <- impar:
			fmt.Printf("Valor %d recebido no canal Impar\n", v)
		case <- quit:
			fmt.Printf("Processo encerrado\n")
			return
		}
	}
}