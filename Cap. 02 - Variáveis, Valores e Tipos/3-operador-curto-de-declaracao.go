package main

import "fmt"

// x := 2 --> Não funciona

func main() {
	// OPERADOR DE DECLARAÇÃO CURTO (:=)

	// Esse operador lembra uma marmota, que é o mascote oficial
	// do Go (The Go Gopher)

	// O Gopher serve para declarar uma variável

	// TIPAGEM AUTOMÁTICA:
	// -------------------
	// Para tipagem automática a gente usa a marmota (Gopher)
	// Neste caso, o operador curto atribui um tipo, com base no
	// valor da variável

	// O Gopher só funciona dentro de um codeblock

	x := 10              // valor numérico
	y := "Olá, bom dia." // string

	fmt.Println(x, y)

	fmt.Println("----------")
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	// float vem de "ponto flutuante" => Flutuante pq a vir-
	// gula pode estar em qualquer lugar entre os caractéres
	// numéricos.

	// OPERADOR DE ATRIBUIÇÃO (=)
	x = 15
	fmt.Printf("x: %v, %T\n\n", x, x)
	// Não pode ser utilizado desta maneira para declaração,
	// apenas para atribuição.

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n\n", z, z)

	// Go posui palavras reservadas que não podem ser uti-
	// lizadas como identificadores.
	// Ver no: "https://go.dev/ref/spec"

	// Toda expressão é feita de operadores e operandos
	// Uma expressão é qualquer coisa que produz m resul-
	// tado, mas que não gera uma ação.

	// declaração --> É uma instrução que forma uma ação

}
