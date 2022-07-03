package main

import "fmt"

func main() {
	var a bool
	fmt.Println(a)
	a = 10 < 20
	fmt.Println(a)

	s := "Hello World @&é"
	fmt.Printf("%v\n%T", s, s)

	sb := []byte(s)
	fmt.Printf("%v\n%T\n", sb, sb)

	for _, i := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", i, i, i, i)
	}

	fmt.Println("")

	for _, i := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", i, i, i, i)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

	fmt.Println("")
	v := 100
	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x\n", v, v, v)

	fmt.Println("")
	const num = 10 // tipo indefinido
	var numInt int16
	var numFloat float64
	numInt = num
	numFloat = num
	fmt.Printf("%v - %T\n", numInt, numInt)
	fmt.Printf("%v - %T\n", numFloat, numFloat)

}
