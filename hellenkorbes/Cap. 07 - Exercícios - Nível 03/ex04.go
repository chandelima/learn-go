/*
### Na prática: exercício #4

- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
- Solução: https://play.golang.org/p/dIbfdiuw0ms
*/

package main

import "fmt"

func main() {

	anoEmQueNasci := 1991
	anoAtual := 2022

	for {
		if anoEmQueNasci <= anoAtual {
			fmt.Println(anoEmQueNasci)
			anoEmQueNasci++
		} else {
			break
		}
	}
}
