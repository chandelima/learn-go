// SLICE - A SURPRESA DO ARRAY SUBJACENTE

package main

import "fmt"

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroslice)

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)
	fmt.Println(segundoslice)
	fmt.Println(primeiroslice)

	/*
		Ao criar um segundo slice a partir do primeiro com append,
		a linguagem manipula o array subjacente do primeiro slice.
		Com isso, os dados do primeiro slice também são manipula-
		dos. O QUE PODE SER UMA PEGADINHA!
		Ao tentar criar um segundo slice, é melhor usar um for!
	*/
}
