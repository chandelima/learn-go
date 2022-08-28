package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - %v", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	p1 := pessoa{
		nome:      "Alexandre",
		sobrenome: "Lima",
	}
	imprimir(p1)

	p2 := produto{
		nome:  "Caneta Azul",
		preco: 1.99,
	}
	imprimir(p2)

	// implementando a interface e aplicando o polimorfismo
	var coisa imprimivel = pessoa{
		nome:      "Roberto",
		sobrenome: "Montezuma",
	}
	imprimir(coisa)

	coisa = produto{
		nome:  "Caderno A4",
		preco: 8.99,
	}
	imprimir(coisa)
}
