package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var count = 0
	var totalGoroutines = 100

	var wg sync.WaitGroup

	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {

			v := count
			runtime.Gosched() // Ativa função Yield
			v++
			count = v

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final do contador: ", count) // Retona valores variados, apesar de 10 functions serem chamadas ao mesmo tempo
}
