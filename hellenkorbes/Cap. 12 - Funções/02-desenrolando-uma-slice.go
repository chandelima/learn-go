/*
Utilizando o operador de enumeração (...) como atributo de uma
função variática a partir de slices.
*/

package main

import "fmt"

func main() {

	sliceSoma := []int{10, 10, 50, 20, 5, 3, 2}
	resultSoma := soma(sliceSoma...)
	fmt.Println(resultSoma)

	// numa funcao variática, você pode passar zero argumentos.
	resultSoma = soma()
	fmt.Println(resultSoma)
}

func soma(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado
}
