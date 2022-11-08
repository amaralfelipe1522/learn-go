package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())       // Quantidade de CPUs antes da execução
	fmt.Println(runtime.NumGoroutine()) // Quantidade de threads tinha antes da execução

	wg.Add(2) // Qtd de Goroutines a serem executadas em paralelo
	go func1()
	go func2()

	fmt.Println(runtime.NumCPU())       // Quantidade de CPUs depois da execução
	fmt.Println(runtime.NumGoroutine()) // Quantidade de threads tinha depoiss da execução

	wg.Wait() // Aguardando todos os Done() retornarem para assim encerrar a função main
}

func func1() {
	defer wg.Done()
	for i := 0; i < 40; i++ {
		fmt.Println("Função 1: ", i)
		time.Sleep(30)
	}
}

func func2() {
	defer wg.Done()
	for i := 0; i < 40; i++ {
		fmt.Println("Função 2: ", i)
		time.Sleep(30)
	}
}
