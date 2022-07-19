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
	exMake = append(exMake, 110) // cria automaticamente outro array com uma capacidade máxima maior
	fmt.Printf("Values: %v\tLength: %v\t Capacity: %v\n", exMake, len(exMake), cap(exMake))

	// Slice Multi dimensionais - Slice of slices
	sliceBi := [][]string{
		{"O", " ", "X"},
		{"O", "X", " "},
		{"X", " ", "O"},
	}
	fmt.Println(sliceBi[0])
	fmt.Println(sliceBi[1])
	fmt.Println(sliceBi[2])

	sliceBi[0][2] = "batata"
	fmt.Println(sliceBi[0])

	// A surpresa do Array Subjacente
	primeiroArray := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroArray)

	segundoArray := append(primeiroArray[:2], primeiroArray[4:]...)
	fmt.Println(segundoArray)

	fmt.Println(primeiroArray) // o segundoArray utilizou o mesmo array subjacente

	fmt.Println(&primeiroArray[0]) // Utilizam o mesmo endereço de memória
	fmt.Println(&segundoArray[0])  // Para evitar isso, recomendasse usar FOR LOOP e copiar cada elemento separadamente do array/slice.. ou sobrescrever a mesma variável

	// Maps - Conjunto de dados com chave e valor, como por exemplo agenda telefonica com nome e telefone
	agendaTelefonica := map[string]int{
		"Celta": 12345678,
		"Carol": 87654321,
	}

	agendaTelefonica["Gopher"] = 101010

	fmt.Println(agendaTelefonica["Celta"])
	fmt.Println(agendaTelefonica["Gopher"])

	searchMap, ok := agendaTelefonica["Kaique"]
	fmt.Printf("Valor: %v\tRetorno Bool: %v\n", searchMap, ok) // Retorna 0 (false) quando não encontra a chave no mapa - comma OK

	// Range em Maps não exibem os valores na mesma ordem
	for key, value := range agendaTelefonica {
		fmt.Printf("Chave: %v\tValor: %v\n", key, value)
	}

	// Deletar um elemento do Map
	delete(agendaTelefonica, "Gopher")
}
