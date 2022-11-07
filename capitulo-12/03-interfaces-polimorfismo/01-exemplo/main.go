package main

import "fmt"

// Interface permite que um valor tenha mais de um tipo, é um conjunto de métodos

// Polimorfismo é poder usar uma mesma função para tipos diferentes

type mensagem struct {
	inputMensagem string
}

type anexo struct {
	inputAnexos []string
}

type bancoDeDados interface {
	insert() bool
}

func (m mensagem) insert() bool {
	fmt.Println("Insert realizado no banco -> ", m.inputMensagem)
	return true
}

func (a anexo) insert() bool {
	for _, value := range a.inputAnexos {
		fmt.Println("Insert realizado no banco -> ", value)
	}
	return true
}

func insertDB(b bancoDeDados) {
	b.insert()
}

func main() {
	mensagem := mensagem{
		"Hello World",
	}

	anexos := anexo{
		[]string{"Foto", "Curriculo"},
	}

	insertDB(mensagem)
	insertDB(anexos)
}
