package main

import "fmt"

func main() {
	// OPERADORES ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25     // era int32
	soma2 := numero1 + numero2 // você não pode operar variáveis de tipos diferentes
	fmt.Println(soma2)

	// OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String1" // Operador de atribuição com tipagem declarada
	variavel2 := "String2"           // Operador de atribuoição com tipagem inferida
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println("\n------------------")
	fmt.Println("Operadores Lógicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// OPERADORES UNÁRIOS
	fmt.Println("\n------------------")
	fmt.Println("Operadores Unários")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 15
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	// OPERADOR TERNÁRIO
	// |-> No GO não existe. Use if/else! Kkkkkkkk
}
