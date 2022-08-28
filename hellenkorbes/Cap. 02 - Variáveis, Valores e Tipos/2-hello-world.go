// O compilador sabe que deve começar a compilar pelo package main
package main

import "fmt"

// A primeira função que o compilador procura para rodar é a função main
// Ela é a porta de entrada e de saída de qualquer aplicação em GO
func main() {
	// Aqui a sua aplicação começa

	fmt.Println("Hello world!", "Oi galera", 100)
	// Link: https://pkg.go.dev/fmt#Println

	numeroDeBytes, erros := fmt.Println("Hello World", "Oi galera", 100)
	fmt.Println(numeroDeBytes, erros)
	// <nil> significa nada (null)
	// A função PrintLn é uma função variática. Isso pode ser identificado
	// a partir dos ... declarado nos parâmetros de sua construção. Vasi-
	// camente significa que você pode passar ilimitados parâmetros daque-
	// le mesmo tipo.

	// Em Go você não pode declarar variáveis e não utilizá-las. Como a
	// linguagem preza por limpeza e performance, ela não deixa você fazer
	// isso.
	// Quando você precisa atribuir uma variável que não vai ser utilizada
	// ao retorno de uma função, usa-se o _ (blank identifier). Desta for-
	// ma o compilador não vai apontar erros. Ex.:
	_, errosEx2 := fmt.Println("Hello World", "Oi galera", 100)
	fmt.Println(errosEx2)

	// VARIÁVEIS (INTRODUÇÃO):
	// "uma variável é um objeto (uma posição na memória) capaz de reter e
	// representar um valor ou expressão."

	// Os tipos primitivos da linguagem GO são:

	x := 16        // números
	y := "strings" // strings
	z := true      // bool (tem esse nome estranho em homenagem a um
	// 				  matemático chamado George Bool)

	fmt.Println(x, y, z)

	// Aqui a sua aplicação termina
}
