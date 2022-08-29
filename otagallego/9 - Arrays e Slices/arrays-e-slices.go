package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	fmt.Println("// Arrays")

	var array1 [5]string
	array1[0] = "Alexandre"
	array1[1] = "Ribeiro"
	array1[2] = "Teixeira"
	array1[3] = "Lima"
	array1[4] = "TOP"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Não deixa com o tamanho dinâmico. Mas infere o tamanho pela quantidade de elementos passados
	fmt.Println(array3)

	fmt.Println("// Slices")

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18) // Adiciona item ao slice
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2) // O slice é uma fatia do array. Ele aponta (ponteiro)

	// Arrays internos
	fmt.Println("----------------")
	fmt.Println("Arrays Internos")
	fmt.Println("----------------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length of slice
	fmt.Println(cap(slice3)) // Capacity of slice

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
