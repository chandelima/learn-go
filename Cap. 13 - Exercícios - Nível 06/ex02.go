/*
### Na prática: exercício #2

- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos
   os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de
  todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
- Solução: https://play.golang.org/p/Tgv3wwxKV-
*/

package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3}

	resultado := soma1(mySlice...)
	fmt.Println(resultado)

	resultado = soma2(mySlice)
	fmt.Println(resultado)
}

func soma1(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func soma2(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
