// SLICES

package main

import "fmt"

func main() {
	// O array tem um número pré-definido de elementos (mais rígido)
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	// No slice você não define um número de elementos na inicialização
	// (mais flexível)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Incrementando slice com append
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	// Atribuindo valor a elemento do slice
	slice2[2] = 12345
	fmt.Println(slice2)
}
