/*
### Na prática: exercício #1

- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
- Solução: https://play.golang.org/p/tpWDzzlO2l
*/

package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}

	for i, num := range intArray {
		fmt.Println(i, num)
	}

	fmt.Printf("%T\n", intArray)
}
