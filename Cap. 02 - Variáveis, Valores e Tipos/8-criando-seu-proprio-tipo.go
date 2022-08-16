package main

import "fmt"

type hotdog int //Criamos um tipo chamado hotdog
// O tipo hotdog possui um tipo adjascente, que é o seu tipo base.
// O tipo base de hotdog é int

var b hotdog

func main() {
	a := 10

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", a)

	// b = a
	// Vc não pode atribuir a um tipo hotdog um tipo int, mesmo
	// que este seja o seu subtipo. Afinal, variáveis em Go são
	// imutáveis.
}
