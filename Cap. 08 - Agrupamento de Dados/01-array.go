// ARRAYS

package main

import "fmt"

// Estou declarando um array de 5 posições do tipo int
var x [5]int

// Estou declarando um array de 5 posições do tipo string
var y [5]string

func main() {
	x[0] = 1 // Estou atribuindo o valor 1 a posição 0 do array
	y[4] = "dez"

	fmt.Println(x[0])
	fmt.Println(x)
	fmt.Println(len(x)) // Função len imprime o tamanho do array
	fmt.Printf("%T\n", x)
	fmt.Println("------------")
	fmt.Println(y[4])
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Printf("%T\n", y)
}
