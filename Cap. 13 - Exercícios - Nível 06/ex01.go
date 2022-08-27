/*
### Na prática: exercício #1

- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
- Solução: https://play.golang.org/p/rxJM5fgI-9
*/
package main

import "fmt"

func main() {
	soma := soma(2, 3)
	fmt.Println(soma)

	result, texto := somaComString(2, 3)

	fmt.Println(texto, result)
}

func soma(x, y int) int {
	return x + y
}

func somaComString(x, y int) (int, string) {
	return x + y, "O resultado da soma é:"
}
