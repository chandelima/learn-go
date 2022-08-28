package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	splits := strings.Split(nomeCompleto, " ")
	p.nome = splits[0]
	p.sobrenome = splits[1]
}

func main() {
	p1 := pessoa{"Alexandre", "Lima"}

	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Leonardo Silva")
	fmt.Println(p1.getNomeCompleto())
}
