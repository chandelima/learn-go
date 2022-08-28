package main

import "fmt"

func main() {

	x := 387

	func(y int) { // parametro da função
		fmt.Println(y, "vezes 8736188 é:")
		fmt.Println(y * 873648)
	}(x) // argumento

	// a função anonima é uma função descartável
}
