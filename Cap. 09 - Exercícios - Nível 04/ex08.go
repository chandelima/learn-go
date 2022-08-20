/*
### Na prática: exercício #8

- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
- Solução: https://play.golang.org/p/nD3TW8VQmH
*/

package main

import "fmt"

func main() {
	fmt.Println("Cap. 9 – Exercícios: Nível #4 – 8 \n")

	chava := map[string][]string{
		"winchester_dean": []string{
			"caçar bichinhos feios", "praticar falsidade ideologica", "dirigir o impala",
		},
		"kuerten_guga": []string{
			"conversar", "Jogar tenis", "fazer propaganda",
		},
		"rodriguez_bender": []string{
			"dobrar qualquer material", "beber cerveja", "dizer morda meu popozão brilhante de metal!",
		},
	}

	for k, v := range chava {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
		fmt.Println()
	}
}
