package main

import "fmt"

func main() {
	x := 12

	// Função como expressão
	y := func(x int) int {

		return x * 873648
	}

	fmt.Println(x, "vezes 873648 é:", y(x))
}
