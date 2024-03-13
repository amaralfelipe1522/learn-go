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

	for i:=0; i < totalGoroutines; i++ {
		go func ()  {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}