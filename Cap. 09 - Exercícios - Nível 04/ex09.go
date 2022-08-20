/*
### Na prática: exercício #9

- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
- Solução: https://play.golang.org/p/3fcvHlt8Lm
*/

package main

import "fmt"

func main() {
	x := map[string][]string{
		"joão_da_silva": {"voar", "codar"},
		"paulo_ferras":  {"codar", "desenhar"},
		"thiago_otashi": {"desenhar", "voar"},
	}
	x["anna_julia"] = []string{"cantar", "dançar"}
	for k1, v1 := range x {
		fmt.Println(k1)
		for _, v2 := range v1 {
			fmt.Println("\t", v2)
		}
	}
}
