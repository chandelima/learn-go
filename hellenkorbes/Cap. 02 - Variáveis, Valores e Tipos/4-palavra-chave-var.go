package main

import "fmt"

// "var" KEYWORD

// O var pode ser utilizado em package level scope. Diferentemente
// do Gopher, que só pode ser utilizado em codeblock level scope.

var x = 2

func main() {
	qualquerCoisa(x)
}

func qualquerCoisa(x int) {
	// y := 1
	var y = 1
	fmt.Println(x)
	fmt.Println(y)
}

// Quando se faz apenas a inicialização de uma variável no package
// level scope, o seu valor só pode ser atribuido dentro de um
// codeblock

// Além dos tipos de dados primitivos (int, string, bool), existem
// os tipos de dados compostos, que são formados pelos usuários, a
// partir de tipos primitivos.
// Ex.: slice, array, struct, map

// O ato de definir/estruturar tipos compostos, chama-se composição
