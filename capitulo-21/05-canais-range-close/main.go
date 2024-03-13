package main

import "fmt"

func main() {
	totalDeAtribuicoes := 10
	canal := make(chan int)

	go loopChannel(totalDeAtribuicoes, canal)

	for v := range canal {
		fmt.Println(v)
	}
}

func loopChannel(total int, s chan<- int) {
	for i:=0; i<total; i++ {
		s <- i
	}
	close(s) // Sem fechar o canal após o loop, ocorrerá deadlock range loop após exibir todos
}