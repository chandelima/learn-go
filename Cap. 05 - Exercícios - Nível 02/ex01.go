/*
Na prática: exercício #1
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
Solução: https://play.golang.org/p/X7qm3aWSa6
*/

package main

import "fmt"

func main() {
	x := 500
	fmt.Printf("%d, %#x, %b", x, x, x)
}
