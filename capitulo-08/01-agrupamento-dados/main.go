package main

import "fmt"

func main() {
	// Array - Inflexicíveis em relação a tipagem e quantidade, nunca é negativo
	array := [2]int{1, 2}
	fmt.Printf("%T\t%v\t%v\t%v\n", array, array[0], array[1], len(array))

	// Slice - Igual ao array, porém não tem um comprimento definido
	slice := []int{1, 2}
	fmt.Printf("%T\t%v\t%v\t%v\n", slice, slice[0], slice[1], len(slice))
	slice = append(slice, 3)
	fmt.Printf("%T\t%v\t%v\t%v\t%v\n", slice, slice[0], slice[1], slice[2], len(slice))

	total := 0

	// For Range
	for _, i := range slice {
		fmt.Println(i)
		total += i
		fmt.Println(total)
	}

	// Slice um Slice
	frutas := []string{"Banana", "Maçã", "Pêra", "Melancia", "Kiwi"}
	vitamina := frutas[0:2]
	for _, fruta := range vitamina {
		fmt.Println(fruta)
	}

	vitamina = frutas[2:]
	for _, fruta := range vitamina {
		fmt.Println(fruta)
	}

	// Removendo um item do array -> Criando uma nova
	frutas = append(frutas[:2], frutas[3:]...)
	fmt.Println(frutas)

	// Make otimiza o custo computacional que o Slice causa, podendo pré estabelecer o comprimento maximo e reservar esse espaço
	exMake := make([]int, 7, 10)
	exMake[0] = 10
	exMake[1] = 20
	exMake[2] = 30
	exMake[3] = 40
	exMake[4] = 50
	exMake[5] = 60
	exMake[6] = 70
	exMake = append(exMake, 80)
	exMake = append(exMake, 90)
	exMake = append(exMake, 100)
	exMake = append(exMake, 110) // aumenta a capacidade máxima automaticamente
	fmt.Printf("Values: %v\tLength: %v\t Capacity: %v\n", exMake, len(exMake), cap(exMake))
}
