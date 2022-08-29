package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// IF INIT: Inicializar variável no próprio IF
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número é menor ou igual a zero e <= 10")
	}

	// fmt.Println(outroNumero) --> Isso não dá certo pq a variável criada no IF INIT só atua no contexto local
}
