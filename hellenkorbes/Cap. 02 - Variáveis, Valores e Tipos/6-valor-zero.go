package main

import "fmt"

var x int // Declaração -> Espaço alocado na memória para ser preenchido

var a int
var b float64
var c string
var d bool

func main() {
	x = 10 // Inicialização e atribuição
	fmt.Printf("%v, %T\n", x, x)
	x = 25 // Atribuição
	fmt.Printf("%v, %T\n", x, x)

	y := 10 // Declaração, inicialização e atribuição
	fmt.Printf("%v, %T\n", y, y)

	// VALOR ZERO
	// Valor que encontra-se em uma variável antes dela ser inicializada
	// pelo usuário

	// Os valores zero padrão da linguagem são:
	// - ints: 0
	// - floats: 0.0
	// - booleans: false
	// - strings: ""
	// - pointers, functions, interfaces, slices, channels, maps: nil

	// Use a marmota (:=) sempre que possível;
	// Use var para package-level scope.

	fmt.Println("Valores Zero:")
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
