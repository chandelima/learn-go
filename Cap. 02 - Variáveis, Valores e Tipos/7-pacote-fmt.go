package main

import "fmt"

func main() {
	fmt.Println("INTERPRETED STRING VS LITERAL STRINGS:")
	fmt.Println("--------------------------------------")

	x := "Amor impossível,\nTão impossível,\nMuito difícil,\n\"Extremamente complicado.\""
	fmt.Println(x)
	// Esta é uma string interpretada. Cada caractére é analisado
	// E pode ser aplicada formatações baseado em como a string é
	// declarada

	y := `Amor impossível,\nTão impossível,\nMuito difícil,\n\"Extremamente complicado.\"`
	fmt.Println(y)
	// Esta é uma string raw (bruta). Ela é impressa da maneira
	// em que é declarada.
}
