package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subt := n1 - n2
	return soma, subt
}

func main() {
	soma := somar(3, 4)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	var resultadoSoma, _ = calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}
