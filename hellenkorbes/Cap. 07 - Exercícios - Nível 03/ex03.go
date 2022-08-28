/*
### Na prática: exercício #3

- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/qnFjiDJzLor
*/

package main

import "fmt"

func main() {

	anoEmQueNasci := 1991
	anoAtual := 2022

	for anoEmQueNasci <= anoAtual {
		fmt.Println(anoEmQueNasci)
		anoEmQueNasci++
	}
}
