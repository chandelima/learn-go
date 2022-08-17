/*
### Na prática: exercício #8

- Crie um programa que utilize a declaração switch, sem switch statement especificado.
- Solução: https://play.golang.org/p/TyIGp4Hi8B
*/

package main

import "fmt"

func main() {
	x := 6

	switch {
	case x < 5:
		fmt.Println("x é menor que 5.")
	case x == 5:
		fmt.Println("x é igual a 5.")
	case x > 5:
		fmt.Println("x é maior que 5.")
	}
}
