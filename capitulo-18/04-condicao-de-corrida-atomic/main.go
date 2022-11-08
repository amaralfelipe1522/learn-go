package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Verificar se há race condition em tempo de execução: go run -race main.go

func main() {

	var count int64 = 0
	var totalGoroutines = 100

	var wg sync.WaitGroup

	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&count, 1) // incrementa valor do contador

			runtime.Gosched() // Ativa função Yield
			time.Sleep(800)

			fmt.Println("Contador: \t ", atomic.LoadInt64(&count))

		}()
	}
	wg.Wait()
	fmt.Println("Valor final do contador: ", count) // Retona valores variados, apesar de 10 functions serem chamadas ao mesmo tempo
}
