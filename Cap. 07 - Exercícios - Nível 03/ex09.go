/*
### Na prática: exercício #9

- Crie um programa que utilize a declaração switch, onde o switch
  statement seja uma variável do tipo string com identificador
  "esporteFavorito".
- Solução: https://play.golang.org/p/4-iTPZwfEz
*/

package main

import "fmt"

func main() {
	esporteFavorito := "espeleoturismo"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("O esporte favorito é futebol")
	case "starcraft":
		fmt.Println("O esporte favorito é starcraft")
	case "espeleoturismo":
		fmt.Println("O esporte favorito é espeleoturismo")
	case "basketball":
		fmt.Println("O esporte favorito é basketball")
	}
}
