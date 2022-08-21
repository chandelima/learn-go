/*
### Na prática: exercício #2

- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	favoritos []string
}

func main() {

	meuMapa := make(map[string]pessoa)

	meuMapa["Lima"] = pessoa{
		nome:      "Alexandre",
		sobrenome: "Lima",
		favoritos: []string{"Chocolate Amargo", "Pistache", "Morango"},
	}

	meuMapa["de Assis"] = pessoa{"Machado", "de Assis", []string{"Manga", "Mangaba", "Cajá"}}

	//fmt.Println(pessoa1)
	//fmt.Println(pessoa2)

	for _, v := range meuMapa {
		fmt.Println(v.nome)
		for _, v := range v.favoritos {
			fmt.Printf("\t- %s\n", v)
		}
		fmt.Println()
	}

}
