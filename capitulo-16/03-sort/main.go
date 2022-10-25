package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"Pera", "Uva", "Maça"}
	fmt.Println(ss)
	fmt.Println(len(ss))

	sort.Strings(ss)
	fmt.Println(ss)

	fmt.Println(sort.SearchStrings(ss, "Uva")) // retornou a posição de Len do slice, e não o index

}
