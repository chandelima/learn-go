package main

import "fmt"

var x int = 10

func main() {
	// EXPLORANDO TIPOS

	// Tipos em Go são estáticos. Uma vez que uma variável é
	// declarada como sendo de um tipo, ela será daquele tipo
	// até o fim da aplicação

	x = 20
	// x = 20.5 --> Não pode, pq seria um float
	fmt.Println(x)

	y := 30
	z := 40.7
	fmt.Printf("y = %v, tipo %T\n", y, y)
	fmt.Printf("z = %v, tipo %T\n", z, z)
}
