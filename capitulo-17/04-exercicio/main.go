package main

import (
	"fmt"
	"sort"
)

type videogame struct {
	Nome          string
	AnoLancamento int
	Jogos         []string
}

type ordenarAnoLancamento []videogame

func (ano ordenarAnoLancamento) Len() int {
	return len(ano)
}

func (ano ordenarAnoLancamento) Less(i, j int) bool {
	return ano[i].AnoLancamento < ano[j].AnoLancamento
}

func (ano ordenarAnoLancamento) Swap(i, j int) {
	ano[i], ano[j] = ano[j], ano[i]
}

type ordenarJogos []string

func (jogo ordenarJogos) Len() int {
	return len(jogo)
}

func (jogo ordenarJogos) Less(i, j int) bool {
	return jogo[i] < jogo[j]
}

func (jogo ordenarJogos) Swap(i, j int) {
	jogo[i], jogo[j] = jogo[j], jogo[i]
}

func main() {
	videogame1 := videogame{
		"Super Nintendo",
		1990,
		[]string{"Super Mario World", "F-Zero"},
	}

	videogame2 := videogame{
		"Mega Drive",
		1989,
		[]string{"Sonic", "Altered Beast"},
	}

	videogame3 := videogame{
		"PlayStation",
		1995,
		[]string{"PES", "Crash"},
	}

	videogames := []videogame{videogame1, videogame2, videogame3}

	fmt.Println(videogames)
	sort.Sort(ordenarAnoLancamento(videogames))
	fmt.Println(videogames)
	for _, videogame := range videogames {
		sort.Strings(videogame.Jogos)
	}
	fmt.Println(videogames)
}
