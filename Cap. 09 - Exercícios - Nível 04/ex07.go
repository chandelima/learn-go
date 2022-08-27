/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
- Solução: https://play.golang.org/p/Gh81-d5tMi
*/

package main

import "fmt"

func main() {
	pessoa := [][]string{
		[]string{"Nome1", "Sobrenome1", "hobbie1"},
		[]string{"Nome2", "Sobrenome2", "hobbie2"},
		[]string{"Nome3", "Sobrenome3", "hobbie3"},
	}
	for _, a := range pessoa {
		fmt.Println(a[0])
		for _, item := range a[1:] {
			fmt.Println("\t item:", item)
		}
	}
}
