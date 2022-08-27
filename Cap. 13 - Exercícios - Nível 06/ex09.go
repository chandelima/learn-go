/*
### Na prática: exercício #9

- Callback: passe uma função como argumento a outra função.
- Solução: https://play.golang.org/p/2epLD_Yyap
*/

package main

import "fmt"

func main() {
	operacao(soma, 100, 10)
	operacao(subt, 100, 10)
	operacao(mult, 100, 10)
	operacao(divi, 100, 10)
}

func operacao(operacao func(a, b int) int, x int, y int) {
	resultado := operacao(x, y)
	fmt.Println(resultado)
}

func soma(x, y int) int {
	return x + y
}

func subt(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func divi(x, y int) int {
	return x / y
}
