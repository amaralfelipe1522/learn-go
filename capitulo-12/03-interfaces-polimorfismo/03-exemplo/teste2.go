package main

import (
	"fmt"
	"strconv"
)

type stringNumber struct {
	number string
}

type intNumber struct {
	number int
}

func (s stringNumber) isPar() bool {
	number, _ := strconv.Atoi(s.number)
	if number%2 != 0 {
		return false
	}
	return true
}

func (i intNumber) isPar() bool {
	if i.number%2 != 0 {
		return false
	}
	return true
}

type typeCalc interface {
	isPar() bool
}

func verifyPar(t typeCalc) {
	fmt.Println(t.isPar())
}

func main() {

	number1 := stringNumber{
		number: "6",
	}

	number2 := intNumber{
		number: 7,
	}

	fmt.Println(number1.isPar())
	fmt.Println(number2.isPar())

	verifyPar(number1)
	verifyPar(number2)
}
