// APPEND: Anexando a uma slice

package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	// Plotando a slice original
	fmt.Println(umaslice)

	// Acrescentando elementos a slice:
	// O comando append recebe além da própria slice, elementos
	// do mesmo tipo da slice a que se quer acrescentar
	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)

	// No entanto, você pode adicionar uma slice a outra slice.
	// Para isso, você deve usar o operator ... (ele "desenola"
	// a slice e retorna os seus elementos).
	// |-> Mesmo que spread operator do JS.
	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)
}
