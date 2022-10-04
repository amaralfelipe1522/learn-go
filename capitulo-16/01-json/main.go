package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// IMPORTANTE - Para que o objeto seja convertido em JSON, os campos devem iniciar com letra maiuscula
// Em GO, tudo que se inicia com letra maiuscula pode ser importada, como um pacote
type Pessoa struct {
	Nome   string   `json:"Nome"`
	Idade  int      `json:"Idade"`
	Cursos []string `json:"Cursos"`
}

func main() {
	// Marshal - Converte um valor para JSON
	pessoa1 := Pessoa{
		"Felipe",
		30,
		[]string{"Golang", "Docker"},
	}
	pessoaJson, err := json.Marshal(pessoa1)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(pessoaJson)
	os.Stdout.Write(pessoaJson)
	fmt.Println("")

	// Unmarshal - Converte um valor para GO
	pessoa2 := []byte(`{"Nome":"Felipe","Idade":30,"Cursos":["Golang","Docker"]}`)

	pessoaGO := Pessoa{}

	err2 := json.Unmarshal(pessoa2, &pessoaGO)
	if err2 != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(pessoaGO)

	//Utilizando função Encoder para converter em JSON
	pessoaEncoder := json.NewEncoder(os.Stdout)
	pessoaEncoder.Encode(pessoa1)

	// Ao contrário de Marshal, o Encoder não depende de uma variável intermediaria para armazenar o resulta da conversão

}
