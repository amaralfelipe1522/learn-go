package main

import (
	"fmt"
	"runtime"
)

// comando para compilar para Windows: GOOS=windows GOARCH=amd64 go build main.go
// comando para compilar para Mac: GOOS=darwin GOARCH=amd64 go build main.go

func main() {
	fmt.Println("Aplicação criada em linux / amd64. Está agora sendo executado em ",runtime.GOOS,"/",runtime.GOARCH)
}