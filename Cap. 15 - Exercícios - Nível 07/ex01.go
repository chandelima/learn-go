/*
### Na prática: exercício #1

- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
- Solução: https://play.golang.org/p/0jVt1yaoFL
*/

package main

import "fmt"

func main() {
	x := 11

	fmt.Println(&x)
}
