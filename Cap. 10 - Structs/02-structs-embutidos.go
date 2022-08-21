package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint
}

type profissional struct {
	pessoa
	titulo  string
	salario uint
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
