package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(3)

	go func1()
	go func2()
	go func3()

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	defer wg.Done()
	for i:=0; i<40; i++ {
		fmt.Printf("Func1 -> %d\n", i)
		time.Sleep(90)
	}
}

func func2() {
	defer wg.Done()
	for i:=0; i<50; i++ {
		fmt.Printf("Func2 -> %d\n", i)
		time.Sleep(90)
	}
}

func func3() {
	defer wg.Done()
	for i:=0; i<60; i++ {
		fmt.Printf("Func3 -> %d\n", i)
		time.Sleep(90)
	}
}