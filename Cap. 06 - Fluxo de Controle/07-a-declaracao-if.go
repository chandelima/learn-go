package main

import "fmt"

func main() {

	x := 100

	if x < 100 {
		fmt.Println("Caso verdadeiro.")
	} else {
		fmt.Println("Caso falso.")
	}

	if y := 101; !(y < 100) { // Você pode declarar uma variável no próprio if
		fmt.Println("Caso verdadeiro.")
	} else {
		fmt.Println("Caso falso.")
	}
}
