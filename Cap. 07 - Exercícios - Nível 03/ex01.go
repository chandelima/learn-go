/*
### Na prática: exercício #1

- Põe na tela: todos os números de 1 a 10000.
- Solução: https://play.golang.org/p/MkdZiDW8SQ
*/

package main

import "fmt"

func main() {
	for x := 1; x <= 10000; x++ {
		fmt.Printf("%v\t", x)
	}
}
