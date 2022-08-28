//Slice: for range

package main

import "fmt"

func main() {
	slice1 := []string{"banana", "maçã", "jaca", "pêssego"}
	for indice, valor := range slice1 {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	fmt.Println("-----------------------------------------------")

	slice1 = append(slice1, "uva")
	for indice, valor := range slice1 {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	fmt.Println("-----------------------------------------------")

	for _, valor := range slice1 {
		fmt.Println("Um dos valores desse slice é:", valor)
	}

	fmt.Println("-----------------------------------------------")

	slice2 := []int{20, 21, 22, 23}
	total := 0
	for _, valor := range slice2 {
		total += valor
	}
	fmt.Println("O valor total é:", total)
}
