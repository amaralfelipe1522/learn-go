package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{32, 54, 12, 2, 67, 45, 90, 1}
	xs := []string{"Pera", "Uva", "Maça", "Amora"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
