/*
### Na prática: exercício #3

- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
- Solução: https://play.golang.org/p/b5Ua2kNN8a
*/

package main

import "fmt"

func main() {
	defer fmt.Println("Última linha")
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}
