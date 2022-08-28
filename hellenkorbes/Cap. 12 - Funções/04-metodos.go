/*
MÉTODOS
Um método é uma função anexada a um tipo
*/

package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint
}

// Este é um método anexo ao tipo pessoa
func (p pessoa) oibomdia() { //É utilizado o receiver pra associar uma função a um tipo
	fmt.Println(p.nome, "diz: bom dia!")
}

func main() {
	alexandre := pessoa{"Alexandre", 30}
	alexandre.oibomdia()
}
