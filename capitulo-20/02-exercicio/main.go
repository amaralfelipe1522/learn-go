package main

import "fmt"

type humanos interface {
	falar()
}

type Pessoa struct {
	Nome string
}

func (p *Pessoa) falar() {
	texto := "Que exercício doido ein"
	fmt.Printf("%s disse: %s\n", p.Nome, texto)
}

func main() {
	pessoa := Pessoa {
		Nome: "Felipe",
	}

	pessoa.falar() // Atalho para (&pessoa).falar()
	(&pessoa).falar() // Maneira 'mais' correta de realizar a chamada
	dizerAlgumaCoisaInterface(&pessoa)
	dizerAlgumaCoisaStruct(pessoa)
}

func dizerAlgumaCoisaInterface(h humanos) {
	h.falar()
}

func dizerAlgumaCoisaStruct(p Pessoa) {
	p.falar()
}