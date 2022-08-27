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
	pessoa1 := pessoa{ // Declaração explícita
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

	pessoa3 := pessoa{"Mauricio", 40} // Declaração implícita

	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 100000000}

	fmt.Println(pessoa1.idade)       // Imprimindo apenas a idade
	fmt.Println(pessoa2.pessoa.nome) // Imprimindo apenas o nome
	fmt.Println(pessoa2.nome)        // Também funciona chamando
	// diretamente campos de um struct embutido, se não houverem
	// conflitos. "promoted fields" é o termo que se dá a isso.
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)
}
