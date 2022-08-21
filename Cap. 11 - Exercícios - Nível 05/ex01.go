/*
### Na prática: exercício #1

- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: https://play.golang.org/p/Pyp7vmTJfY
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	favoritos []string
}

func favoritos(people pessoa) {
	fmt.Printf("%s %s:\n", people.nome, people.sobrenome)
	for _, v := range people.favoritos {
		fmt.Println("\t", v)
	}
}

func main() {
	pessoa1 := pessoa{
		nome:      "Alexandre",
		sobrenome: "Lima",
		favoritos: []string{"Chocolate Amargo", "Pistache", "Morango"},
	}

	pessoa2 := pessoa{"Machado", "de Assis", []string{"Manga", "Mangaba", "Cajá"}}

	//fmt.Println(pessoa1)
	//fmt.Println(pessoa2)

	for _, v := range pessoa1.favoritos {
		fmt.Println(v)
	}

	favoritos(pessoa1)
	favoritos(pessoa2)
}
