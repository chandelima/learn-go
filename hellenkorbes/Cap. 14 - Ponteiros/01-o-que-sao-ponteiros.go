package main

import "fmt"

func main() {
	// Ponteiro é uma variável que armazena endereço de memória

	x := 10
	fmt.Println(x)
	/*
		fmt.Println(&x) // Notação "&" mostra o endereço da memória

		y := &x         // Atribuí o endereço de memória de x a y
		fmt.Println(*y) // * faz o de-reference do endereço e mostra o valor
	*/

	y := &x

	*y++

	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Println(x, y)
}
