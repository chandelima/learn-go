/*
### Na prática: exercício #2

- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
    - ==
    - !=
    - <=
    - <
    - >=
    - >
- Demonstre estes valores.
- Solução: https://play.golang.org/p/BMYEch6_s8
*/

package main

import "fmt"

func main() {
	a := (1 == 2)
	b := (1 != 2)
	c := (1 <= 1)
	d := (1 < 2)
	e := (1 >= 2)
	f := (1 > 1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
