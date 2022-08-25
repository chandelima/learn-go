/*
Funções servem para abstrair funcionalidades.

Estrutura de uma função:
func (receiver) identifier(parameters) (returns) { code }

*/

package main

import "fmt"

func main() {

	// basica()

	// argumento("manhã")
	// argumento("tarde")
	// argumento("noite")

	// resultSoma, lenSoma := soma(10, 10)
	// fmt.Println(resultSoma, lenSoma)

	mensagem, resultado := soma2("O resultado da soma é:", 10, 10, 10, 10, 10)
	fmt.Println(mensagem, resultado)
}

func basica() {
	fmt.Println("Oi, bom dia.")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia.")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde.")
	} else {
		fmt.Println("Oi, boa noite.")
	}
}

// Função variática (número qualquer de argumentos)
// e com dois retornos.
func soma(x ...int) (int, int) {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado, len(x)
}

// Você pode criar funções variáticas com parâmetros
// de outros tipos. Neste caso, o parâmetro variáti-
// co sempre deve ser declarado por último.
func soma2(s string, x ...int) (string, int) {
	resultado := 0
	for _, v := range x {
		resultado += v
	}

	return s, resultado
}
