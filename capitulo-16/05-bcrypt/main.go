package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// go get -u golang.org/x/crypto/bcrypt

// Cost:
// MinCost int 4,
// MaxCost int 31,
// DefaultCost int 10,

func main() {
	senha := "Golang@1522"
	fmt.Println("Senha em claro: ", senha)

	sliceOfBytes, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hash da senha: ", string(sliceOfBytes))

	errCompare := bcrypt.CompareHashAndPassword(sliceOfBytes, []byte("Golang@1523"))
	if errCompare != nil {
		log.Fatal(errCompare)
	}
	fmt.Println("Senha confirmada")
}
