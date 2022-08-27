/*
### Na prática: exercício #7

- Atribua uma função a uma variável.
- Chame essa função.
- Solução: https://play.golang.org/p/RMHLL3N5Ww
*/

package main

import "fmt"

func main() {
	x := func(a, b int) {
		res := a + b
		fmt.Println(res)
	}

	x(1, 2)
}
