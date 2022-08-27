/*
### Na prática: exercício #8

- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
- Solução: https://play.golang.org/p/A74rufv6Rs
*/

package main

import "fmt"

func main() {
	soma := operacoes()
	fmt.Println(soma(2, 3))
}

func operacoes() func(x, y int) int {
	return func(x, y int) int {
		return x + y
	}
}
