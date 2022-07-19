package main

import "fmt"

func main() {
	// Exercício 1
	array := [5]int{
		1, 2, 3, 4, 5,
	}
	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Printf("Tipo do Array: %T\n", array)

	// Exercício 2
	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("Tipo do slice: %T\n", slice)

	// Exercício 3

	// 1º ao 3º item
	fmt.Println(slice[:3])
	// 5º ao último item
	fmt.Println(slice[4:])
	// 2º ao 7º item
	fmt.Println(slice[1:7])
	// 3º ao penultimo item
	fmt.Println(slice[2:9])
	// 3º ao penultimo item usando len
	fmt.Println(slice[2 : len(slice)-1])

	// Exercício 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	// Exercício 5
	x2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y2 := append(x2[:3], x2[6:]...)
	fmt.Println(y2)

	// Exercício 6
	estados := make([]string, 26, 27)
	fmt.Printf("Length: %v\tCapacity: %v\n", len(estados), cap(estados))

	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Printf("*Length: %v\tCapacity: %v\n", len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}

	// Exercício 6
	pessoas := [][]string{
		{"Felipe", "Amaral", "Videogames"},
		{"Celta", "Tolansky", "Carros"},
		{"Carol", "Carles", "Violão"},
	}

	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %v\tSobrenome: %v\tHobby favorito: %v\n", pessoa[0], pessoa[1], pessoa[2])
	}

	// Exercício 7
	pessoasMap := map[string][]string{
		"Felipe_Amaral":  {"Videogames", "Animes"},
		"Celta_Tolansky": {"Carros", "Corrida"},
		"Carol_Carles":   {"Violão", "Volei"},
	}

	pessoasMap["Gabriel_Fosco"] = []string{"Videogames", "Animes"}

	delete(pessoasMap, "Felipe_Amaral")

	for nome_sobrenome, hobbies := range pessoasMap {
		fmt.Println(nome_sobrenome)
		for _, hobby := range hobbies {
			fmt.Println(hobby)
		}
	}
}
