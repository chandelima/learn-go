/*
### Na prática: exercício #10

- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
- Solução:
*/

package main

import (
	"fmt"
)

func main() {
	a := map[string][]string{
		"Abrantes_Wagner":  []string{"Filmes", "Programmar"},
		"Abrantes_Beatriz": []string{"Ru Paul", "Beber Chá"},
		"Andrade_Maria":    []string{"Reclamar", "Igreja"},
	}

	//fmt.Println(a)

	delete(a, "Abrantes_Beatriz")

	for k, v := range a {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
