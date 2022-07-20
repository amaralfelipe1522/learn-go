package main

import "fmt"

type funcionario struct {
	nome     string
	idade    int
	endereco endereco
	status   bool
}

type endereco struct {
	rua string
	cep int
}

func main() {
	// Struct
	enderecoFuncionarioA := endereco{
		"Rua Masuzo",
		123456,
	}
	fmt.Println(enderecoFuncionarioA.rua, enderecoFuncionarioA.cep)

	// Struct embutido = Struct dentro de outra Struct
	funcionarioA := funcionario{
		"Felipe",
		31,
		enderecoFuncionarioA,
		true,
	}
	fmt.Println(funcionarioA.nome, funcionarioA.idade, funcionarioA.endereco.rua, funcionarioA.status)

	// Struct anonimos = Não são reutilizáveis
	head := struct {
		nome  string
		idade int
	}{
		"Nazaré",
		40,
	}
	fmt.Println(head.nome)
}
