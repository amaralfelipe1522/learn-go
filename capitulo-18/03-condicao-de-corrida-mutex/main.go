package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Verificar se há race condition em tempo de execução: go run -race main.go

func main() {

	var count = 0
	var totalGoroutines = 100

	var wg sync.WaitGroup

	var mutex sync.Mutex

	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			defer wg.Done()

			mutex.Lock()

			v := count
			runtime.Gosched() // Ativa função Yield
			time.Sleep(800)
			v++
			count = v

			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final do contador: ", count) // Retona valores variados, apesar de 10 functions serem chamadas ao mesmo tempo
}
