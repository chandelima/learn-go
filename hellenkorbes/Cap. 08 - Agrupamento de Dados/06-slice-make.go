// SLICE - MAKE

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 6)
	fmt.Println(slice, len(slice), cap(slice))
	// Neste momento o tamanho do slice passa a ser 6, e a capacity
	// do array interno passa a ser 10. Por padrão o Go sempre do-
	// bra a capacidade do array interno. No entanto, essa operação
	// pode ser custosa, pois quando o array interno dobra de capa-
	// cidade, é na verdade realocado um novo espaço na memória, e
	// criado um novo array.

	// Neste contexto, pode ser interessante utilizar a função make
	slice2 := make([]int, 5, 10) // <- Último argumento da função é
	// justamente a capacidade que se quer atribuir ao array inter-
	// no a ser criado.
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice2[0], slice2[1], slice2[2], slice2[3], slice2[4] = 1, 2, 3, 4, 5
	slice2 = append(slice2, 6, 7, 8, 9, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))
	// Note que neste caso, apesar do slice ter sido incrementado,
	// a capacidade do array interno não foi dobrada. Isso pq a
	// mesma foi pré-definida anteriormente com a declaração a
	// partir da função make

}
