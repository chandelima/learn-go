package main

import "fmt"

func main() {

	/*
		Structs anonimos são estruturas "temporárias", onde você inicia-
		liza a estrutura de um struct em uma variável e logo em seguida
		a popula. É feito para ser utilizado uma única vez. Exemplo a-
		baixo.
	*/

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}

	fmt.Println(x)
}
