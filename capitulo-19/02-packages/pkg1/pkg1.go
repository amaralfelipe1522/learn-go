package main

import (
	"fmt"
	"runtime"
)

func Pkg1() {
	fmt.Println("Arquitetura atual: ", runtime.GOARCH)
}

func Somar(x, y int) int {
	return x + y
}