package main

import "fmt"

func main() {
	x := 0

	// FOR SÓ COM CONDIÇÃO:
	// for x < 10 {
	// 	fmt.Println("x é menor que 10. x =", x)
	// 	x++
	// }

	/*
		Em Go não há while. No entanto, além do for tradicional
		(em que se declara valor inicial, condição, incremento)
		você pode declarar só a condição e iterar uma variável
		fora do escopo do looping. Agindo como um while.
	*/

	// FOR SEM NENHUM ARGUMENTO
	// Neste caso, os controles devem ser lançados dentro do
	// statement (ou não, se for o caso)
	for {
		if x < 10 {
			x++
			fmt.Println("x < 10")
		} else {
			break
		}
	}
}
