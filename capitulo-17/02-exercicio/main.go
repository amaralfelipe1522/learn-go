package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type videogame struct {
	Nome          string
	AnoLancamento int
	Jogos         []string
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

	videogames := []videogame{videogame1, videogame2}

	fmt.Println(videogames)

	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(videogames)

	// ou

	json.NewEncoder(os.Stdout).Encode(videogames)

}
