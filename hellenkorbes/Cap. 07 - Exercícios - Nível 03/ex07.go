/*
### Na prática: exercício #7

- Utilizando a solução anterior, adicione as opções else if e else.
- Solução: https://play.golang.org/p/_cO7E-yL0o
*/

package main

import "fmt"

func main() {
	if x := 10; x < 5 {
		fmt.Println("x é menor que 5.")
	} else if x == 5 {
		fmt.Println("x é igual a 5.")
	} else {
		fmt.Println("x é maior que 5.")
	}
}
