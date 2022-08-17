/*
### Na prática: exercício #6

- Crie um programa que demonstre o funcionamento da declaração if.
- Solução: https://play.golang.org/p/OGPgTJZbiy
*/

package main

import "fmt"

func main() {
	if x := 2; x < 5 {
		fmt.Println("x é menor que 5.")
	}
}
