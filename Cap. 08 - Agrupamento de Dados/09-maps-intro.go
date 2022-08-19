// MAPS - INTRODUÇÃO

package main

import "fmt"

func main() {
	// O Map é uma estrutura chave valor da linguagem. Ex:
	// Map que representa o telefone dos amigos.
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	// Imprimindo na tela o map inteiro
	fmt.Println(amigos)

	// Imprimindo na tela um único elemento do map
	fmt.Println(amigos["alfredo"], "\n")

	// No entanto, repare no comportamento do map ao ser impres-
	// so o valor de chave não declarada
	fmt.Println(amigos["romario"])
	// Neste caso, é plotado o valor zero do tipo atribuído.
	// Para saber se realmente trata-se de um valor declarado ou
	// do próprio valor zero, usa-se o "comma ok idiom"

	// Comma ok idiom: É um segundo retorno do elemento do map,
	// Que retorna um booleano que informa se aquela chave foi
	// atribuída, ex.:

	// romario, ok := amigos["romario"]

	// fmt.Println(romario, ok) // Neste caso, a chave nao foi
	// atribuída

	// Exemplo de validação de comma ok com if
	if romario, ok := amigos["romario"]; ok {
		fmt.Println(romario)
	} else {
		fmt.Println("Valor não atribuído.")
	}
}
